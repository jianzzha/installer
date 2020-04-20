#!/bin/bash

parse_args() {
   WRITE_USB=${WRITE_USB:-true}

   USAGE="Usage: $0 [options]
Options:
    -d USB_DEV     Install to the given usb device, alternatively
                   specify env var USB_DEV
    -i IGNITION    The URL or path to the given Ignition config,
                   alternatively specify env var IGNITION_URL
    -o INSTALLER   The URL or path to the installer image,
                   alternatively specify env var INSTALLER_IMAGE
    -m METAL_IMG   The URL or path to the Metal image, 
                   alternatively specify env var METAL_IMAGE
    -b BLOCK_DEV   The block device on the remote node,
                   alternatively specify env var BLOCK_DEV,
		   default to /dev/sda
    -h             This.

This tool builds a USB key that can be used to install remote nodes when pxe boot is not feasible.
"
    while getopts "d:i:o:m:b:h" OPTION
    do
        case $OPTION in
            d) USB_DEV="$OPTARG" ;;
            i) IGNITION_URL="$OPTARG" ;;
	    o) INSTALLER_IMAGE="$OPTARG" ;;
            m) METAL_IMAGE="$OPTARG" ;;
            b) BLOCK_DEV="$OPTARG" ;;
            h) echo "$USAGE"; exit;;
            *) echo "$USAGE"; exit 1;;
        esac
    done

    if [[ -z "${USB_DEV}" && ${WRITE_USB} == "true" ]]; then
	echo "usb device must be specified unless WRITE_USB is set to false!"
	echo "$USAGE"
	exit 1
    fi
    # normalize device name
    USB_DEV=$(echo ${USB_DEV} | sed -r 's|/dev/||')
    USB_DEV=/dev/${USB_DEV}

    if [[ -z "${IGNITION_URL}" ]]; then
	echo "ignition url must be specified!"
	echo "$USAGE"
	exit 1
    fi
  
    if [[ -z "${INSTALLER_IMAGE}" ]]; then
        echo "installer image url must be specified!"
	echo "$USAGE"
	exit 1
    fi

    if [[ -z "${METAL_IMAGE}" ]]; then
	echo "metal image url must be specified!"
	echo "$USAGE"
	exit 1
    fi

    BLOCK_DEV=${BLOCK_DEV:-sda}

    # by default flash usb key and clean up intermediate file folders after flash
    # otherwise don't flash and leave the intermediate files for debug
}


main() {
    parse_args $@

    # test device make sure it is present
    if ! lsblk ${USB_DEV} > /dev/null && [[ ${WRITE_USB} == "true" ]]; then
	echo "usb device must be inserted unless WRITE_USB is set to false!"
	exit 1
    fi

    for link in ${IGNITION_URL} ${INSTALLER_IMAGE} ${METAL_IMAGE}; do
	if ! curl -LsIf ${link} >/dev/null 2>&1; then
            echo "invalid link: ${link}"
	    exit 1
	fi
    done

    /bin/rm -rf _data
    mkdir -p _data
    pushd _data
    mkdir -p _extra/extra
    echo "Downloading ignition file ..."
    if ! curl -L -s -o _extra/extra/node.ign ${IGNITION_URL} >/dev/null 2>&1; then
	echo "failed to get ${IGNITION_URL}!"
	exit 1
    fi
    echo "Downloading installer ..."
    if ! curl -L -s -o installer.iso ${INSTALLER_IMAGE} >/dev/null 2>&1; then
	echo "failed to get ${INSTALLER_IMAGE}!"
        exit 1
    fi
    echo "Downloading metal image ..."
    if ! curl -L -s -o _extra/extra/rhcos.raw.gz ${METAL_IMAGE} >/dev/null 2>&1; then
	echo "failed to get ${METAL_IMAGE}!"
	exit 1
    fi

    # setup booting parameter
    /bin/rm -rf _rhcos_custom
    mkdir _rhcos_custom
    mkdir _iso_mount

    echo "mounting installer iso ..."
    mount -t iso9660 -o loop installer.iso _iso_mount 

    echo "copy out installer files ..."
    pushd _iso_mount 
    tar cf - . | (cd ../_rhcos_custom && tar xfp -)
    popd

    echo "update installer parameters ..."
    pushd _rhcos_custom 
    sed -i -r "s|^(  append initrd)=.*|\1=/images/initramfs.img,/images/initramfsExtra nomodeset ip=dhcp rd.neednet=1 coreos.inst=yes console=tty0 console=ttyS0 coreos.inst.ignition_url=file:///extra/node.ign coreos.inst.image_url=file:///extra/rhcos.raw.gz coreos.inst.install_dev=${BLOCK_DEV} ramdisk_size=8388608|g" isolinux/isolinux.cfg
    popd

    echo "building extra initramfs ..."
    pushd _extra
    find . | sed 's/^[.]\///' | cpio -o -H newc --no-absolute-filenames > ../_rhcos_custom/images/initramfsExtra
    popd
   
    echo "rebuilding installer image ..."
    pushd _rhcos_custom
    mkisofs -o rhcos_custom_usb.iso -b isolinux/isolinux.bin -c isolinux/boot.cat  -no-emul-boot -boot-load-size 4 -boot-info-table -R -J -V "RHCOS custom installer" .
    popd

    if [[ "${WRITE_USB}" == "false" ]]; then
	echo "WRITE_USB=${WRITE_USB}; complete without writing usb key"
	popd
	exit 0
    fi

    echo "writing usb disk ..."
    pushd _rhcos_custom
    dd bs=4M if=rhcos_custom_usb.iso of=${USB_DEV} status=progress oflag=sync
    popd

    echo "cleaning up ..."
    umount _iso_mount
    popd
    /bin/rm -f _data

    echo "completed"
}

main $@

