help:
	@echo "help"
	@echo
	@echo "manuf: Pulls latest /etc/manuf from code.wireshark.org"

manuf:
	wget -O manuf "https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf"
