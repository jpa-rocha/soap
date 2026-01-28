{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nix-filter.url = "github:numtide/nix-filter";
    nixpkgs-stable.url = "github:nixos/nixpkgs/nixos-25.05";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    pre-commit-hooks.url = "github:cachix/git-hooks.nix";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs =
    { self, ... }@inputs:
    let
      eachSystem = inputs.flake-utils.lib.eachDefaultSystem;
    in
    eachSystem (
      system:
      let
        pkgs = inputs.nixpkgs.legacyPackages.${system};
        lib = pkgs.lib;

        inherit (lib) makeBinPath;
        filterOwn = inputs.nix-filter.lib;
        go = pkgs.go_1_24;
        goBuild = pkgs.buildGo124Module;

        gowsdl = pkgs.callPackage ./nix/gowsdl.nix { inherit go goBuild makeBinPath; };

        golangci-lint = pkgs.callPackage ./nix/golangci-lint.nix { inherit go goBuild makeBinPath; };

        scripts = import ./nix/scripts.nix { inherit pkgs go; };
        soap = pkgs.callPackage ./nix/soap.nix { inherit goBuild filterOwn; };

        devInputs = with pkgs; [
          git
          git-chglog
          gofumpt
          golines
          gopls
          gowsdl
          sops
          scripts.lint
          scripts.test
          scripts.tidy
          soap
          go
          golangci-lint
          golangci-lint-langserver
        ];

        treefmtEval = inputs.treefmt-nix.lib.evalModule pkgs ./nix/treefmt.nix;

        preCommit = import ./nix/pre-commit.nix {
          inherit golangci-lint;
          gofumpt = pkgs.gofumpt;
          nixfmt-rfc-style = pkgs.nixfmt-rfc-style;
          prettier = pkgs.nodePackages.prettier;
          trufflehog = pkgs.trufflehog;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = devInputs;
          inherit (self.checks.${system}.pre-commit-check) shellHook;
        };

        checks = {
          formatting = treefmtEval.config.build.check inputs.self;
          pre-commit-check = inputs.pre-commit-hooks.lib.${system}.run preCommit;
        };

        formatter = treefmtEval.config.build.wrapper;

        packages.soap = soap;
      }
    );
}
