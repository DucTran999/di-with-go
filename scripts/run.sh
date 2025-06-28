#!/bin/bash

cat <<"EOF"
  ____ ___    ____  ___  
 |  _ \_ _|  / ___|/ _ \ 
 | | | | |  | |  _| | | |
 | |_| | |  | |_| | |_| |
 |____/___|  \____|\___/ 

EOF

todos=()

function show_menu() {
    echo "=== DI Menu ==="
    echo "1) Example DI manual"
    echo "2) Example DI with wire"
    echo "3) Exit"
    echo -n "Select option: "
}

show_menu
read option
case $option in
1)
    go run ./cmd/manual/
    ;;
2)
    go run ./cmd/wire/
    ;;
3)
    echo "👋 Goodbye!"
    exit 0
    ;;
*)
    echo "❌ Invalid option."
    ;;
esac
echo ""
