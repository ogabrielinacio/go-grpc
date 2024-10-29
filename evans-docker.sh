docker run --network="host"  -it --rm -v "$(pwd):/mount:ro" \
    ghcr.io/ktr0731/evans:latest \
      --path ./proto \
      --proto course_category.proto \
      --host  127.0.0.1 \
      --port 5051 \
      repl
