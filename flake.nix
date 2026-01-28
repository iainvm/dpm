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
    nixpkgs,
    flake-utils,
    ...
  }: let
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
    ];
  in
    {
      packages.x86_64-linux.default = pkgs.buildGoModule {
        pname = "dpm";
        version = "0.1.0";
        src = ./.;
        vendorHash = "sha256-PeAC4Pf7YksxUJOFpVTxdGnmgEZ/IdazttCg452eEXQ=";
        buildInputs = build_deps;
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
