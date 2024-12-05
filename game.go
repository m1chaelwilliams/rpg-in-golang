package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"rpg-tutorial/scenes"
)

type Game struct {
	sceneMap      map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.GameSceneId:  scenes.NewGameScene(),
		scenes.StartSceneId: scenes.NewStartScene(),
		scenes.PauseSceneId: scenes.NewPauseScene(),
	}
	activeSceneId := scenes.StartSceneId
	sceneMap[activeSceneId].FirstLoad()
	return &Game{
		sceneMap,
		activeSceneId,
	}
}

func (g *Game) Update() error {
	nextSceneId := g.sceneMap[g.activeSceneId].Update()
	// switched scenes
	if nextSceneId == scenes.ExitSceneId {
		g.sceneMap[g.activeSceneId].OnExit()
		return ebiten.Termination
	}
	if nextSceneId != g.activeSceneId {
		nextScene := g.sceneMap[nextSceneId]
		// if not loaded? then load in
		if !nextScene.IsLoaded() {
			nextScene.FirstLoad()
		}
		nextScene.OnEnter()
		g.sceneMap[g.activeSceneId].OnExit()
	}
	g.activeSceneId = nextSceneId
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneMap[g.activeSceneId].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
