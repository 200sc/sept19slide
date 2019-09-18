package main

import (
	"fmt"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/render/mod"
	"github.com/oakmound/oak/shape"
	"golang.org/x/image/colornames"
)

const (
	width  = 1920
	height = 1080
)

var (
	Express28  = show.FontSize(28)(show.Express)
	Gnuolane28 = show.FontSize(28)(show.Gnuolane)
	Libel28    = show.FontSize(28)(show.Libel)

	RLibel28 = show.FontColor(colornames.Blue)(Libel28)

	Express44  = show.FontSize(44)(show.Express)
	Gnuolane44 = show.FontSize(44)(show.Gnuolane)
	Libel44    = show.FontSize(44)(show.Libel)

	Express72  = show.FontSize(72)(show.Express)
	Gnuolane72 = show.FontSize(72)(show.Gnuolane)
	Libel72    = show.FontSize(72)(show.Libel)
)

func main() {

	show.SetDims(width, height)
	show.SetTitleFont(Gnuolane72)

	bz1, _ := shape.BezierCurve(
		width/15, height/5,
		width/15, height/15,
		width/5, height/15)

	bz2, _ := shape.BezierCurve(
		width-(width/15), height/5,
		width-(width/15), height/15,
		width-(width/5), height/15)

	bz3, _ := shape.BezierCurve(
		width/15, height-(height/5),
		width/15, height-(height/15),
		width/5, height-(height/15))

	bz4, _ := shape.BezierCurve(
		width-(width/15), height-(height/5),
		width-(width/15), height-(height/15),
		width-(width/5), height-(height/15))

	bkg := render.NewCompositeM(
		render.NewColorBox(width, height, colornames.Indianred),
		render.BezierThickLine(bz1, colornames.White, 1),
		render.BezierThickLine(bz2, colornames.White, 1),
		render.BezierThickLine(bz3, colornames.White, 1),
		render.BezierThickLine(bz4, colornames.White, 1),
	)

	oak.LoadingR = bkg

	setups := []slideSetup{
		intro,
		engines,
		examples,
		philo,
		fullGames,
		conclusion,
	}

	total := 0

	for _, setup := range setups {
		total += setup.len
	}

	fmt.Println("Total slides", total)

	sslides := static.NewSlideSet(total,
		static.Background(bkg),
	)

	nextStart := 0

	for _, setup := range setups {
		setup.add(nextStart, sslides)
		nextStart += setup.len
	}

	oak.SetupConfig.Screen = oak.Screen{
		Width:  width,
		Height: height,
	}
	oak.SetupConfig.FrameRate = 30
	oak.SetupConfig.DrawFrameRate = 30

	slides := make([]show.Slide, len(sslides))
	for i, s := range sslides {
		slides[i] = s
	}
	show.AddNumberShortcuts(len(slides))
	show.Start(slides...)
}

type slideSetup struct {
	add func(int, []*static.Slide)
	len int
}

var (
	intro = slideSetup{
		addIntro,
		5,
	}
)

func addIntro(i int, sslides []*static.Slide) {
	// Intro: three slides
	// Title
	sslides[i].Append(
		show.Title("Oak: a pure Go game engine"),
		show.TxtAt(Gnuolane44, "Patrick Stephen", .5, .6),
	)
	// Thanks everybody for coming to this talk. I'm going to be talking about
	// design patterns, philosophies, and generally useful tricks for
	// developing video games in Go.

	sslides[i+1].Append(show.Header("Who Am I"))
	sslides[i+1].Append(
		show.TxtSetAt(Gnuolane44, 0.5, 0.63, 0.0, 0.07,
			"Go Architect at Ardan Labs",
			"Maintainer / Programmer of Oak",
			"github.com/200sc  github.com/oakmound/oak",
			"patrick.d.stephen@gmail.com",
			"oakmoundstudio@gmail.com",
		)...,
	)
	// My name is Patrick Stephen, I'm currently working at Ardan Labs
	// subcontracted out as a Go Architect. I'm one of two primary maintainers
	// of oak's source code, Oak being the subject of this talk.
	//
	// If you have any questions that don't get answered in or after
	// this talk, feel free to send those questions either to me
	// personally or to our team's email, or if it applies, feel free
	// to raise an issue on the repository.

	sslides[i+2].Append(show.Header("What is Oak?"))
	sslides[i+2].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Game engine in Go",
			"- Targets zero non-Go dependencies",
			"- Primarily 2D",
		)...,
	)

	sslides[i+3].Append(show.Header("This Talk is Not About..."))
	sslides[i+3].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Optimizing Go",
			"- 3D Graphics in Go",
			"- Mobile Games in Go",
			"- Tutorializing Oak",
		)...,
	)

	// And just to get this out of the way, as you will probably
	// note from the games I just showed, we aren't going to be
	// talking about 3D games here or really performance intensive
	// games, or games for non-desktop platforms, just because,
	// while we haven't ignored these things I don't have
	// any revolutionary breakthroughs to share about them right now.

	sslides[i+4].Append(show.Header("Topics"))
	sslides[i+4].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Game Engines in Go",
			"- Small Example Programs",
			"- Design Philosophy",
			"- Full, Released Games",
		)...,
	)

	// What we will talk about, is why Go is particularly useful for
	// developing games, the philosophy behind our engine and development
	// strategy, and then some interesting use cases for applying
	// design patterns that Go makes easy with particle generation,
	// artificial intelligence, and level construction.
}

var (
	examples = slideSetup{
		addExamples,
		1,
	}
)

func addExamples(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Examples"))
}

var (
	fullGames = slideSetup{
		addFullGames,
		1,
	}
)

func addFullGames(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Full Games"))
}

var (
	engines = slideSetup{
		addEngines,
		7,
	}
)

func addEngines(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Other Game Engines in Go"))
	sslides[i+1].Append(show.Header("Abandoned Engines"))
	sslides[i+1].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Azul3D",
			"- GarageEngine",
			"- glop",
		)...,
	)
	sslides[i+1].Append(show.ImageAt("azul_logo.PNG", .7, .5, mod.Scale(.25, .25)))

	sslides[i+2].Append(show.Header("Termloop"))
	sslides[i+2].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- In-Terminal Games (2D)",
			"- First Release: Never",
		)...,
	)
	sslides[i+2].Append(show.ImageAt("termloop_maze.PNG", .7, .5, mod.Scale(.25, .25)))

	sslides[i+3].Append(show.Header("Engo"))
	sslides[i+3].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Primarily 2D",
			"- First Release: 1.0.0 in May 2018",
			"- Last Release: 1.0.4 in April 2019",
		)...,
	)
	sslides[i+3].Append(show.ImageAt("engo_logo.PNG", .7, .5, mod.Scale(1, 1)))

	sslides[i+4].Append(show.Header("G3N"))
	sslides[i+4].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Primarily 3D",
			"- First Release: 0.1.0 in Sept 2019",
			"- Last Release: 0.1.0 in Sept 2019",
		)...,
	)
	sslides[i+4].Append(show.ImageAt("g3n_logo.PNG", .7, .5, mod.Scale(.5, .5)))
	sslides[i+5].Append(show.Header("Pixel"))
	sslides[i+5].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Primarily 2D",
			"- First Release: 0.5.0 in April 2017",
			"- Last Release: 0.8.0 in Oct 2018",
		)...,
	)
	sslides[i+5].Append(show.ImageAt("Pixel.PNG", .7, .5, mod.Scale(.5, .5)))
	sslides[i+6].Append(show.Header("Ebiten"))
	sslides[i+6].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Primarily 2D",
			"- Supports Mobile/WASM",
			"- First Release: 1.0.0 in June 2016",
			"- Last Release: 1.10.0-alpha in Sept 2019",
		)...,
	)
	sslides[i+6].Append(show.ImageAt("ebitenFlappy.PNG", .7, .5, mod.Scale(.5, .5)))
}

var (
	philo = slideSetup{
		addPhilo,
		7,
	}
)

func addPhilo(i int, sslides []*static.Slide) {
	// Philosophy, engine discussion
	sslides[i].Append(show.Title("Design Philosophy"))
	sslides[i+1].Append(show.Header("Design Philosophy"))
	sslides[i+1].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- No non-Go dependencies",
			"- Ease / Terseness of API",
			"- If it's useful and generic, put it in the engine",
		)...,
	)

	// That brings us to our design philosophy in Oak.
	// First, if we have a non-Go dependency, we also have an issue to
	// replace that non-Go dependency ASAP. Right now we have just one.
	//
	// The motivation for having few dependencies isn't just so we can
	// feel confident that all of our platforms are supported, but also
	// making the engine easy to use. Most libraries in Go can be used
	// with 'go get', and we want the same thing here-- a developer
	// should be able to go get oak and immediately start working with it.
	//
	// After that, we want our API to be easy to use and small. Part of our
	// motivation to start building Oak was that other game engines at the
	// time took 500 lines to draw a cube or 400 lines to write Pong. Maybe
	// at their core, those problems do take that many lines, but a lot of that
	// code we can write for you (and also for us, so we don't have to keep
	// re-writing it).
	//
	// In line with this, we follow a rule where if we have to rewrite something
	// more than once for a game or for a package of the engine, that probably
	// means that should be its own package and feature the engine provides.
	// This does go against the go proverbs-- we do not follow the idea that
	// a little copying is better than a little dependency, so long as we
	// treat that dependency as part of the larger, engine dependency.

	sslides[i+2].Append(show.Header("Update Loops and Functions"))
	sslides[i+2].Append(
		show.Image("updateCode1.PNG", .27, .4),
		show.Image("updateCode3.PNG", .57, .4),
	)
	//
	// Some game engines model their exposed API as a loop--
	// stick all your logic inside update()
	//
	// In larger projects, this leads directly to an explicit splitting up of that
	// loop into at least two parts-- update all entities, then
	// draw all entities.
	//
	// The combining of these elements into one loop causes
	// a major problem-- tying the rate at which entities update themselves
	// to the rate at which entities are drawn. This leads to inflexible
	// engines, and in large projects you'll have to do something to work around
	// this, or if you hard lock your draw rate modders will post funny videos
	// of your physics breaking when they try to fix your frame rate.
	//
	// Oak handles this loop for you, and splits it into two loops, one for
	// drawing elements and one for logical frame updating.
	//
	sslides[i+3].Append(show.Header("Update Loops and Functions"))
	sslides[i+3].Append(
		show.Image("updateCode2.PNG", .27, .4),
		show.Image("updateCode3.PNG", .57, .4),
	)
	//
	// Another pattern used, in parallel with the Update Loop,
	// is the Update Function. Give every entity in your game the
	// Upate() function, and then your game logic is handled by calling Update()
	// on everything. At a glance, this works very well in Go because your entities
	// all fit into this single-function interface, but in games with a lot of
	// entities you'll end up with a lot of entities that don't need to do
	// anything on each frame.
	//
	// The engine needs to provide a way to handle game objects that don't
	// need to be updated as well as those that do, and separating these into
	// two groups explicitly makes the engine less extensible. Oak uses an
	// event handler for this instead, where each entity that wants to use
	// an update function binds that function to their entity id once.
	//
	sslides[i+4].Append(show.Header("Useful Packages"))
	sslides[i+4].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- oak/alg/intgeom, oak/alg/floatgeom",
			"- oak/joystick",
			"- oak/render/particle",
			"- oak/entities/x/btn",
		)...,
	)
	//
	// These are some of the less obvious useful packages we've taken
	// from games or sub-packages and built into their own package--
	//
	// intgeom and floatgeom should be self explanatory-- we and every
	// other Go package continually redefine X,Y and X,Y,Z points of
	// integers and floats, and we needed to stop redoing that work.
	//
	sslides[i+5].Append(show.Header("oak/alg"))
	sslides[i+5].Append(show.ImageAt("degToRad.PNG", .3, .5, mod.Scale(1.25, 1.25)))
	sslides[i+5].Append(show.ImageAt("chooseX.PNG", .6, .5, mod.Scale(1.25, 1.25)))
	//
	// in alg, we store things like rounding and selection algorithms.
	// We found that we really needed to pick a random element from
	// a list of weighted floats a lot, so we split it off here.
	//
	sslides[i+6].Append(show.Header("oak/entities/x/btn"))
	sslides[i+6].Append(show.ImageAt("buttonCode.PNG", .5, .5, mod.Scale(1.25, 1.25)))
}

var (
	conclusion = slideSetup{
		addConclusion,
		2,
	}
)

func addConclusion(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Header("Thanks To"))
	sslides[i].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Nate Fudenberg, John Ficklin",
			"- Contributors on Github",
			"- You, Audience",
		)...,
	)

	// And I'll end by thanking the other people in Oakmound
	// for working with me on our engine, those who've tried
	// out the engine and raised issues or PRs, and all of you
	// for listening.

	sslides[i+1].Append(show.Title("Questions"))

	// ?
}
