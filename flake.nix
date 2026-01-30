{
  description = "DPM";

  inputs = {
    nixpkgs = {
      url = "github:NixOS/nixpkgs";
    };

    flake-utils = {
      url = "github:numtide/flake-utils";
    };
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }: let
    version = "0.1.0";
    pkgs = import nixpkgs {system = "x86_64-linux";}; # Specify system type
    build_deps = with pkgs; [
      go_1_24
    ];
    dev_deps = with pkgs; [
      # Git
      git

      # Nix LSP
      nixd
      # Nix Formatter
      alejandra

      # Taskfile
      go-task

      # Go
      gopls
      delve
      gotools
      golangci-lint
    ];
  in
    {
      packages.x86_64-linux.default = pkgs.buildGoModule {
        pname = "dpm";
        version = version;
        src = ./.;
        vendorHash = "sha256-PeAC4Pf7YksxUJOFpVTxdGnmgEZ/IdazttCg452eEXQ=";
        # vendorHash = "";
        buildInputs = build_deps;
        ldflags = [
          "-s -w"
          "-X main.version=${version}"
          "-X main.buildDate=${self.lastModifiedDate}"
        ];
      };
    }
    // flake-utils.lib.eachDefaultSystem (system: {
      devShells = {
        default = pkgs.mkShell {
          shellHook = ''
            export GOPRIVATE="github.com/iainvm"
            export GONOPROXY="github.com/iainvm"
          '';

          # Required for debugging Go
          hardeningDisable = ["fortify"];

          packages = build_deps ++ dev_deps;
        };
      };
    });
}
