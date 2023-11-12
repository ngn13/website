#!/bin/bash
echo -n "Enter API URL: "
read url
cat > /usr/bin/admin_script << EOF
#!/bin/sh
API=$url python3 $(pwd)/admin.py \$1
EOF
chmod +x /usr/bin/admin_script
