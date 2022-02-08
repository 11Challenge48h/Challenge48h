<script setup></script>

<template>
  <div class="fond">
    <head>
      <link
        href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
        rel="stylesheet"
      />
      <title>STAR WARS DATA PAGE</title>
    </head>
    <div class="row">
      <div class="profil">
        <h2>GUENARD Juliette</h2>
        <h3>B1 Computer Science Ynov Nantes</h3>
        <a href="https://www.linkedin.com/in/juliette-guenard-2a11ba220/"
          ><img class="linkedin" src="https://zupimages.net/up/22/06/5t0a.png"
        /></a>
        <p>
          <img class="people" src="https://zupimages.net/up/22/06/4a5w.png" />
        </p>
      </div>
      <div class="profil">
        <h2>GRELLIER Mattéo</h2>
        <h3>B2 Computer Science Ynov Nantes</h3>
        <a href="https://www.linkedin.com/in/matteo-grellier-54a9101a1/"
          ><img class="linkedin" src="https://zupimages.net/up/22/06/5t0a.png"
        /></a>
        <p>
          <img class="people" src="https://zupimages.net/up/22/06/fzpk.png" />
        </p>
      </div>
      <div class="profil">
        <h2>LO Tchu</h2>
        <h3>B3 Computer Science, dev, Ynov Nantes</h3>
        <a href="https://www.linkedin.com/in/tchu-lo-b07b79226/"
          ><img class="linkedin" src="https://zupimages.net/up/22/06/5t0a.png"
        /></a>
        <p>
          <img class="people" src="https://zupimages.net/up/22/06/cygd.png" />
        </p>
      </div>
    </div>
    <div class="row">
      <div class="profil">
        <h2>BOURRY Nathan</h2>
        <h3>B1 Computer Science Ynov Nantes</h3>
        <a href="https://www.linkedin.com/in/nathan-bourry-6a2709208/"
          ><img class="linkedin" src="https://zupimages.net/up/22/06/5t0a.png"
        /></a>
        <p>
          <img class="people" src="https://zupimages.net/up/22/06/9kgt.png" />
        </p>
      </div>
      <div class="profil">
        <h2>CLAVIER Arnaud</h2>
        <h3>B2 Computer Science Ynov Nantes</h3>
        <a href="https://www.linkedin.com/in/arnaud-clavier-6594111b0/"
          ><img class="linkedin" src="https://zupimages.net/up/22/06/5t0a.png"
        /></a>
        <p>
          <img class="people" src="https://zupimages.net/up/22/06/9a2s.png" />
        </p>
      </div>
      <div class="profil">
        <h2>NEIL Romain</h2>
        <h3>B3 Computer Science, dev, Ynov Nantes</h3>
        <a href="https://www.linkedin.com/in/romain-neil/"
          ><img class="linkedin" src="https://zupimages.net/up/22/06/5t0a.png"
        /></a>
        <p>
          <img class="people" src="https://zupimages.net/up/22/06/jx46.png" />
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fond {
  background: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 0.5),
    rgba(241, 225, 0, 0.671) 50%
  );
}
.linkedin {
  max-width: 5%;
  max-height: 5%;
  margin: auto;
  margin-bottom: 10px;
}
.row {
  margin-top: 35px;
  display: flex;
  justify-content: space-between;
}

.profil {
  flex-basis: 25%;
  background: black;
  border-radius: 10px;
  padding: 3px 20px;
  box-sizing: border-box;
  margin: 50px;
  box-shadow: 15px 15px 15px;
}

.profil h2 {
  text-align: center;
  margin: 10px 0;
  color: white;
  font-size: xx-large;
}

.profil h3 {
  text-align: center;
  font-weight: 600;
  margin: 10px 0;
  color: white;
  font-family: "Prompt";
}

.profil img {
  width: 100%;
  height: 100%;
  display: block;
}
</style>
