{
  system,
  inputs,
  ...
}: let
  pkgs = inputs.nixpkgs.legacyPackages.${system};
in {
  default = pkgs.mkShell {
    packages = with pkgs; [
      # Git
      git

      # Nix LSP
      nixd
      # Nix Formatter
      alejandra

      # Taskfile
      go-task

      # Go
      go_1_24
      gopls
      delve
      gotools
    ];
  };
}
