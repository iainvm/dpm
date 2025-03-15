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
  } @ inputs: let
    pkgs = import nixpkgs {system = "x86_64-linux";}; # Specify system type
  in
    {
      packages.x86_64-linux.default = pkgs.buildGoModule {
        pname = "dpm";
        version = "0.1.0";
        src = ./.;
        vendorHash = "sha256-iR5t7Blx+hOGlph5L+C8VWLUo9dnzTKlzP/BJlCcKso=";
        buildInputs = [
          pkgs.go
        ];
      };
    }
    // flake-utils.lib.eachDefaultSystem (system: {
      devShells = import ./devShell/configuration.nix {
        inherit inputs system;
      };
    });
}
