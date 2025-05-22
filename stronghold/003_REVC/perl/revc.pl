#!/usr/bin/env perl

sub ReverseComplement {
  my $seq = $_;
  $seq = reverse($seq);
  $seq =~ tr/ACGT/TGCA/;
  return $seq;
}

while (<>) {
  my $seq = $_;
  $seq = chomp($seq);
  $seq = ReverseComplement($seq);
  print("$seq\n");
}
