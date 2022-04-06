{pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {

    nativeBuildInputs = [
        # Go
        pkgs.buildPackages.go
        pkgs.buildPackages.gotools
        pkgs.buildPackages.gopls
        pkgs.buildPackages.go-outline
        pkgs.buildPackages.gocode
        pkgs.buildPackages.gopkgs
        pkgs.buildPackages.gocode-gomod
        pkgs.buildPackages.godef
        pkgs.buildPackages.golint
    ];
}
