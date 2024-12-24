#!/bin/bash

# Print the SFS ASCII art banner
cat << "EOF"
 _____                           ______ _ _        _____ _                _
/  ___|                          |  ___(_) |      /  ___| |              (_)
\ `--.  ___  ___ _   _ _ __ ___  | |_   _| | ___  \ `--.| |__   __ _ _ __ _ _ __   __ _
 `--. \/ _ \/ __| | | | '__/ _ \ |  _| | | |/ _ \  `--. \ '_ \ / _` | '__| | '_ \ / _` |
/\__/ /  __/ (__| |_| | | |  __/ | |   | | |  __/ /\__/ / | | | (_| | |  | | | | | (_| |
\____/ \___|\___|\__,_|_|  \___| \_|   |_|_|\___| \____/|_| |_|\__,_|_|  |_|_| |_|\__, |
                                                                                   __/ |
                                                                                  |___/
EOF

# Output message indicating the SFS executable is being run
echo "Running the SFS executable..."

# Execute the SFS binary
$HOME/sfs
