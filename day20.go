package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type set struct {
	x int
	y int
	z int
}

type particle struct {
	pos set
	vel set
	acc set
}

func moveParticle(in *particle) {
	in.vel.x += in.acc.x
	in.vel.y += in.acc.y
	in.vel.z += in.acc.z
	in.pos.x += in.vel.x
	in.pos.y += in.vel.y
	in.pos.z += in.vel.z
}

func getParticleDistance(in particle) int {
	return 0 //abs(in.pos.x) + abs(in.pos.y) + abs(in.pos.z)
}

func advent20A(test string) int {
	var particles []particle
	re := regexp.MustCompile(`(-?)\d+`)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		p := re.FindAllString(s, -1)
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		pz, _ := strconv.Atoi(p[2])
		vx, _ := strconv.Atoi(p[3])
		vy, _ := strconv.Atoi(p[4])
		vz, _ := strconv.Atoi(p[5])
		ax, _ := strconv.Atoi(p[6])
		ay, _ := strconv.Atoi(p[7])
		az, _ := strconv.Atoi(p[8])
		particles = append(particles, particle{set{px, py, pz}, set{vx, vy, vz}, set{ax, ay, az}})
	}
	for i := 0; i < 1000; i++ {
		for p, particle := range particles {
			moveParticle(&particle)
			particles[p] = particle
		}
	}
	closest := 0
	closestDistance := 999999999999999999
	for p, particle := range particles {
		distance := getParticleDistance(particle)
		if distance < closestDistance {
			closestDistance = distance
			closest = p
		}
	}

	return closest
}

func annilihate(particles []particle) []particle {
	var aftermath []particle
	for i, particleA := range particles {
		collision := false
		for j, particleB := range particles {
			if i == j {
				continue
			}
			if particleA.pos.x == particleB.pos.x && particleA.pos.y == particleB.pos.y && particleA.pos.z == particleB.pos.z {
				collision = true
				break
			}
		}
		if !collision {
			aftermath = append(aftermath, particleA)
		}
	}
	return aftermath
}

func advent20B(test string) int {
	var particles []particle
	re := regexp.MustCompile(`(-?)\d+`)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		p := re.FindAllString(s, -1)
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		pz, _ := strconv.Atoi(p[2])
		vx, _ := strconv.Atoi(p[3])
		vy, _ := strconv.Atoi(p[4])
		vz, _ := strconv.Atoi(p[5])
		ax, _ := strconv.Atoi(p[6])
		ay, _ := strconv.Atoi(p[7])
		az, _ := strconv.Atoi(p[8])
		particles = append(particles, particle{set{px, py, pz}, set{vx, vy, vz}, set{ax, ay, az}})
	}
	for i := 0; i < 1000; i++ {
		particles = annilihate(particles)
		for p, particle := range particles {
			moveParticle(&particle)
			particles[p] = particle
		}
	}

	return len(particles)
}
