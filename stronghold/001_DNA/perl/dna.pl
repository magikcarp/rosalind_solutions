#!/usr/bin/env perl

sub CountLetters {
  my $seq = $_[0];
  my %counts;

  map {$counts{$_}++} split(//, $_);
  
  return %counts;
}

while (<>) {
  my $seq = $_;
  $seq = uc(chomp($seq));
  my %nuc_counts = CountLetters($seq);

  print(join(" ", (
    $nuc_counts{'A'},
    $nuc_counts{'C'},
    $nuc_counts{'G'},
    $nuc_counts{'T'}
  )), "\n");
}
