package main

import (
  "errors"
)

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

type MapData struct {
  Field         [20][79]byte
  LadsRemaining int
  Level         int
  Score         int
  Bonustime     int
}

type XY struct {
  x, y int
}

//
// This loads one of the playfields/maps into memory and also
// returns the coordinates of the initial Lad, and an array of
// coordinates where the dispensers are.
//
func LoadMap(n int) (MapData, XY, []XY, error) {
  var x, y int
  var m MapData
  var lad XY
  var dispensers []XY

  if n > len(levels) {
    return m, lad, dispensers, errors.New("Level out of range")
  }

  // Prepare the field to be loaded with a new level
  for y = 0; y < 20; y++ {
    for x = 0; x < 79; x++ {
      m.Field[y][x] = ' '
    }
  }

  for y = 0; y < len(levels[n].layout); y++ {
    for x = 0; x < len(levels[n].layout[y]); x++ {
      switch levels[n].layout[y][x] {
      case 'p':
        // The lad will be put there by the rendered, so no need to have it on the map
        lad.x = x
        lad.y = y
      case 'V':
        m.Field[y][x] = 'V'
        dispensers = append(dispensers, XY{x: x, y: y})
      case '.': // TODO - handle the rubber balls
        m.Field[y][x] = '.'
      default:
        m.Field[y][x] = levels[n].layout[y][x]
      }
    }
  }

  return m, lad, dispensers, nil
}
