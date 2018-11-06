#!/usr/bin/bash
echo "Starting Automated VM Configuration..."
echo "Creating new user, 'munir_test'..."
sudo useradd munir_test
echo "User 'munir_test', created..."
sudo mkdir /home/munir_test/.ssh
#sudo chmod 777 /home/munir_test/.ssh
#sudo chmod 777 /home/munir_test/.ssh/authorized_keys
touch authorized_keys
echo "This is the Authorized Keys File for the Authorized Keys of the User" > authorized_keys
sudo cp ~/authorized_keys /home/munir_test/.ssh
sudo chown munir_test:munir_test /home/munir_test/.ssh
sudo chown munir_test:munir_test /home/munir_test/.ssh/authorized_keys
#sudo su - $username -c 'command goes here'
