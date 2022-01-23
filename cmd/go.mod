module aptitude

go 1.17

require aptitude/internal/number v1.0.0

replace aptitude/internal/number => ../internal/number

require aptitude/internal/perception v1.0.0

replace aptitude/internal/perception => ../internal/perception

require aptitude/internal/reason v1.0.0

replace aptitude/internal/reason => ../internal/reason

require aptitude/internal/spatial v1.0.0

replace aptitude/internal/spatial => ../internal/spatial

replace aptitude/internal/utility => ../internal/utility

require aptitude/internal/utility v1.0.0

require aptitude/internal/word v1.0.0

require fyne.io/fyne/v2 v2.1.2 // indirect

replace aptitude/internal/word => ../internal/word
