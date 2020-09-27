sudo groupadd uinput
sudo usermod -a -G uinput esimov
sudo udevadm control --reload-rules
echo "SUBSYSTEM==\"misc\", KERNEL==\"uinput\", GROUP=\"uinput\", MODE=\"0660\"" | sudo tee /etc/udev/rules.d/uinput.rules
echo uinput | sudo tee /etc/modules-load.d/uinput.conf