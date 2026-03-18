module transcription
    implicit none
contains
    subroutine transcribe(seq, seq_len)
        implicit none
        integer :: seq_len
        character(len=seq_len) :: seq
        integer :: i
        character(len=1) :: thymine = "T"
        character(len=1) :: uracil = "U"

        do i=1,len(seq)
            if (seq(i:i) == thymine) then
                seq(i:i) = uracil
            end if
        end do
    end subroutine
end module

program rna
    use transcription
    implicit none
    integer, parameter :: s_len = 1000
    character (len=s_len) :: sequence
    read *,sequence
    call transcribe(sequence, s_len)
    print *,sequence
end program
