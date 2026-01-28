{ goBuild, filterOwn, ... }:
let
  srcFilter = filterOwn {
    root = ./.;
    include = [
      ./cmd
      ./go.mod
      ./go.sum
      ./internal
      ./main.go
      ./vendor
    ];
  };
in
goBuild {
  name = "soap";
  src = srcFilter;
  doCheck = false;
  vendorHash = null;
  ldflags = [
    "-s"
    "-w"
    "-extldflags '-static'"
  ];
  env.CGO_ENABLED = 0;
}
