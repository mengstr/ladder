package main

import "errors"

// p - The place where the lad starts.
// V - Der Dispenser. Der Rocks roll out of it to squash you flat.
// * - Der Eaters. They eat the Der Rocks but oddly do not harm you in the slightest
// = - Floor. You walk on it.
// H - Ladder. You climb it.
// | - Wall. You can't walk through it. You're not a ghost....yet.
// . - Rubber Ball. It's very bouncy. This difference is, it bounces you.
// $ - Treasure. The lad must get here to finish the level.
// & - Gold Statue. Money!Money!Money!Money!Money!
// ^ - Fire. Turns you into extra crispy bacon.
// - - Disposable Floor. Well, you can walk on it once.

var levels = []struct {
  name   string
  layout []string
}{
  {"Easy Street",
    []string{
      "                                       V                 $                     ",
      "                                                         H                     ",
      "                H                                        H                     ",
      "       =========H==================================================            ",
      "                H                                                              ",
      "                H                                                              ",
      "                H          H                             H                     ",
      "================H==========H==================   ========H=====================",
      "                &          H                             H          |       |  ",
      "                                                         H         Easy Street ",
      "                H                                        H                     ",
      "       =========H==========H=========  =======================                 ",
      "                H                                                              ",
      "                H                                                              ",
      "                H                                        H                     ",
      "======================== ====================== =========H==============       ",
      "                                                         H                     ",
      "                                                         H                     ",
      "*    p                                                   H                    *",
      "===============================================================================",
    },
  },

  {"Long Island",
    []string{
      "                                                                          $    ",
      "                                                                   &      H    ",
      "    H       |V                                                     V|     H    ",
      "====H======================= ========================= ======================  ",
      "    H                                                                          ",
      "    H                                                                          ",
      "    H                    & |                         . .                  H    ",
      "========================== ======  =================== ===================H==  ",
      "                                                                          H    ",
      "                                  |                                       H    ",
      "    H                             |                 .  .                  H    ",
      "====H=====================   ======  ================  ======================  ",
      "    H                                                                          ",
      "    H                      |                                                   ",
      "    H                      |                        .   .                 H    ",
      "=========================  ========    ==============   ==================H==  ",
      "                                                                          H    ",
      "==============                      |                                     H    ",
      " Long Island |   p         *        |                 *                   H    ",
      "===============================================================================",
    },
  },

  {"Ghost Town",
    []string{
      "                            V               V           V               $      ",
      "                                                                       $$$     ",
      "     p    H                                                    H      $$$$$   H",
      "==========H===                                                =H==============H",
      "          H                                                    H              H",
      "          H                              &                     H              H",
      "     ==============   ====     =    ======    =   ====    =====H=====         H",
      "    G              ^^^    ^^^^^ ^^^^      ^^^^ ^^^    ^^^                     $",
      "    h                                                                 |        ",
      "    o     |                     H                             &       |        ",
      "    s     ======================H============================== ===========    ",
      "    t        &                  H                                              ",
      "                                H                                              ",
      "              |                 H                 H                   H        ",
      "    T         ==================H=================H===================H======= ",
      "    o                                             H                   H        ",
      "    w                                                                 H        ",
      "    n                           ^                                     H        ",
      "*                              ^^^                                    H       *",
      "===============================================================================",
    },
  },

  {"Tunnel Vision",
    []string{
      "                                            V                       V          ",
      "                                                                               ",
      "     H             H                         |                H                ",
      "=====H=====--======H==========================     ===----====H===========     ",
      "     H             H                |&&                       H                ",
      "     H             H                ==================        H                ",
      "     H             H                       tunnel  H          H                ",
      "     H           =======---===----=================H=         H           H    ",
      "     H         |                           vision  H          H           H    ",
      "     H         =========---&      -----============H          H           H    ",
      "     H           H                                 H |        H           H    ",
      "     H           H=========----===----================        H  ==============",
      "                 H                                        &   H                ",
      "                 H                                        |   H                ",
      "====---====      H                                        |   H                ",
      "|         |    ================---===---===================   H                ",
      "|   ===   |                                                   H        H    p  ",
      "|    $    |                                                   H     ===H=======",
      "|*  $$$  *|   *                *       *                     *H       *H       ",
      "===============================================================================",
    },
  },

  {"Point of No Return",
    []string{
      "         $                                                                     ",
      "         H                                                   V                 ",
      "         H                                                                     ",
      "         HHHHHHHHHHHHH     .HHHHHHHHHHHHHH                          H    p     ",
      "         &                   V           H                        ==H==========",
      "                                         H                          H          ",
      "   H                                     H        .                 H          ",
      "===H==============-----------============H====                      H          ",
      "   H                                                      H         H          ",
      "   H                                                 =====H==============      ",
      "   H                                     H                H                    ",
      "   H              &..^^^.....^..^ . ^^   H==---------     H                    ",
      "   H         ============================H    &           H             H      ",
      "   H         ===      ===      ===       H    ---------=================H======",
      "   H                                     H                              H      ",
      "   H                          &          H          &                   H      ",
      "   ==========-------------------------=======----------===================     ",
      "                                                                               ",
      "^^^*         ^^^^^^^^^^^^^^^^^^^^^^^^^*     *^^^^^^^^^^*Point of No Return*^^^^",
      "===============================================================================",
    },
  },

  {"Bug City",
    []string{
      "        Bug City             HHHHHHHH                          V               ",
      "                           HHH      HHH                                        ",
      "   H                                          >mmmmmmmm                        ",
      "   H===============                   ====================          H          ",
      "   H              |=====         /          V                  =====H==========",
      "   H                             /                                  H          ",
      "   H                                        | $                     H          ",
      "   H           H                            | H                     H          ",
      "   H       ====H=======          p          |&H    H                H          ",
      "   H           H             ======================H           ======          ",
      "   H           H      &|                           H                    H      ",
      "   H           H      &|                    H      H               =====H====  ",
      "===H===&       H       =====================H      H                    H      ",
      "               H                            H      H                    H      ",
      "               H                            H      &                    H      ",
      "         ======H===   =======    H    <>    &                           H      ",
      "                                 H==========       =====     =     ============",
      "     i                           H                                             ",
      "*                                H                                            *",
      "===============================================================================",
    },
  },

  {"GangLand",
    []string{
      "                    =Gang Land=                             V                  ",
      "                   ==      _  ==                                      .        ",
      "      p    H        |  [] |_| |                  &                    .  H     ",
      "===========H        |     |_| |       H         ===   ===================H     ",
      "      V    H        =============     H======                            H     ",
      "           H                          H                     &            H     ",
      "           H                          H                |    |            H     ",
      "    H      H        ^^^&&^^^ & ^  ^^^ H           H    |    =============H     ",
      "    H======H   =======================H===========H=====          &      H     ",
      "    H                                 H           H    |         &&&     H     ",
      "    H                                 H           H    |        &&&&&    H     ",
      "    H                                 H           H    |    =============H     ",
      "              =====------=================        H    |       $     $         ",
      "                                         |        H    |      $$$   $$$        ",
      "====------===                            |        H    |     $$$$$ $$$$$       ",
      "            |       =                    | =============    ============       ",
      "            |       $                     ^          &                         ",
      "            |^^^^^^^^^^^^^^      $ ^              ======                       ",
      "*                   .      &   ^ H*^                    ^  ^       ^^^^^^^^^^^^",
      "===============================================================================",
    },
  },
}

type Coords struct {
  x int
  y int
}

type MapData struct {
  lad           Coords
  dispensers    []Coords
  ladsRemaining int
  level         int
  score         int
  bonustime     int
  field         [20][79]byte
}

//
//
func loadMap(n int) (MapData, error) {
  if n > len(levels) {
    return MapData{}, errors.New("Level out of range")
  }
  var i, j int
  var m MapData

  for i = 0; i < 20; i++ {
    for j = 0; j < 79; j++ {
      m.field[i][j] = '.'
    }
  }

  for i = 0; i < len(levels[n].layout); i++ {
    //    fmt.Println(levels[n].layout[i])
    for j = 0; j < len(levels[n].layout[i]); j++ {
      switch levels[n].layout[i][j] {
      case 'p':
        m.field[i][j] = 'p'
        m.lad = Coords{i, j}
      case 'V':
        m.field[i][j] = 'V'
        m.dispensers = append(m.dispensers, Coords{i, j})
      case '.': // TODO - handle the rubber balls
        m.field[i][j] = '.'
      default:
        m.field[i][j] = levels[n].layout[i][j]
      }
    }
  }

  return m, nil
}

// p - The place where the lad starts.
// V - Der Dispenser. Der Rocks roll out of it to squash you flat.
// * - Der Eaters. They eat the Der Rocks but oddly do not harm you in the slightest
// . - Rubber Ball. It's very bouncy. This difference is, it bounces you.
