# Dependencies

* pkg-config
* libvirt-dev

# Install as a service

* Check/Update Xen URI in govirt-api.service

* Run these commands:

        % go build -o govirt-api main.go
        % sudo cp govirt-api /usr/local/bin 
        % sudo cp govirt-api.service /lib/systemd/system/
        % sudo systemctl enable govirt-api
        % sudo systemctl start govirt-api