.PHONY: all install-mage

all: install-mage

# Install Mage
install-mage:
	@echo "Downloading Mage..."
	@wget -q https://github.com/magefile/mage/releases/download/v1.9.0/mage_1.9.0_Linux-64bit.tar.gz -O /tmp/mage.tar.gz
	@echo "Extracting Mage..."
	@tar -xf /tmp/mage.tar.gz -C /tmp/
	@echo "Ensuring the directory exists and is empty..."
	@mkdir -p $GOPATH/bin
	@rm -f $GOPATH/bin/mage
	@echo "Moving Mage to the GOPATH/bin directory..."
	@mv /tmp/mage $GOPATH/bin/mage
	@chmod +x $GOPATH/bin/mage
	@echo "Mage installation complete."
