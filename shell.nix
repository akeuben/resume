{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = [ pkgs.texlive.combined.scheme-full pkgs.gnumake ];
}

