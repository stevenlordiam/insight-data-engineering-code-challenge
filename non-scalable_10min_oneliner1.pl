use warnings;use strict;undef $/;my %count=();while($_=<>){foreach (split){s/\W//g;s/^\d*$//;if(length){$count{lc()}++;}}};foreach my $word (sort keys(%count)){printf "%-15s\t%d\n", $word, $count{$word};}
