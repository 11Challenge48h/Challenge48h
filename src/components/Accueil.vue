<script setup>

</script>

<template>
 <head>
    <title>STAR WARS DATA PAGE</title>
  </head>

  <body>
    <section class="main">
      <section class="page1" id="page1">
        <div class="intro">
          <h1>
            STAR WARS DATA<br /><span>Que la force soit avec toi !</span
            ><br />Menu
          </h1>
          <p class="text">
            Hangman-web consists in creating and running a server, in which it
            will be possible to use a web GUI (graphical user interface) version
            of our last project, hangman-classic. We had to use a go module to
            call our functions from hangman-classic.
            <br />
            We needed to create a private repository with the name hangman-web
          </p>
          <button>PLAY NOW</button>
        </div>

        <div>
          <img
            class="photopendu"
            src="https://zupimages.net/up/22/06/ns5k.png"
          />
        </div>
      </section>
    </section>
    <section class="page2" id="page2">
      <h2>THE TEAM</h2>
      <p>Let's introduce ourselves</p>
    </section>
    <section class="page3" id="page3">
      <h2>Nous trouver</h2>
      <p>How to play Hangman ?</p>
      <br />
      <h3>
        You have 10 attempts to guess the random word.<br />
        Every time you make a mistake, José gets closer to death.<br />
        Don't let José hang himself!<br />
        Good luck!
      </h3>
      <div class="row"></div>
    </section>
  </body>
  {{ end }}
</template>

<style scoped>
.main {
  padding: 0 15%; /* marges éspacement des bords*/
  background: linear-gradient(
    /*fond de la première page dégradé*/ to bottom,
    /* on peut aussi mettre to top, to right, to left*/ rgba(255, 255, 255, 0.5),
    rgba(241, 225, 0, 0.671) 50%
  );
  height: 100vh; /*configure pour que la première page remplisse l'écran*/
  display: flex;
}

.page1 {
  color: #fff;
  display: flex;
  margin: 20% 0;
}

.intro .text {
  padding-bottom: 100px;
  letter-spacing: 1.2px;
  line-height: 30px;
}

.intro h1 {
  font-size: 50px;
  letter-spacing: 2px;
  margin-bottom: 35px;
}

.intro > button {
  margin-left: 25%;
  width: 160px;
  height: 50px;
  border: 3px solid rgb(0, 0, 0);
  background-color: transparent;
  font-size: 30px;
  cursor: pointer;
  transition: 0.8s ease;
  color: rgb(0, 0, 0);
  border-radius: 10%;
}

.intro > button:hover {
  color: #000000;
  background-color: white;
}

.intro span {
  color: #000000;
  font-size: 65px;
}

.form h2 {
  text-align: center;
  color: #000000;
  font-size: 22px;
  padding: 8px;
  text-transform: uppercase;
}

h3 {
  text-align: center;
}

.form > .btn a {
  text-decoration: none;
  transition: 0.3s ease;
}

.form .link {
  font-size: 17px;
  padding-top: 20px;
  text-align: center;
}

.form .link a {
  text-decoration: none;
  color: rgba(218, 207, 55, 0.671);
}

.page2,
.page4,
.page3 {
  border-radius: 5px;
  padding: 10% 15%;
}

.page3 {
  background: #753e1914;
}

.page2 h2,
.page4 h2,
.page3 h2 {
  margin-bottom: 25px;
  font-size: 30px;
  font-weight: bold;
  color: #000;
  text-align: center;
}

.page2 p,
.page4 p,
.page3 p {
  text-align: center;
}

.col {
  flex-basis: 40%;
  background: rgba(218, 207, 55, 0.671);
  border-radius: 10px;
  padding: 30px 12px;
  box-sizing: border-box;
  width: 10px;
}

.col h3 {
  text-align: center;
  font-weight: 600;
  margin: 10px 0;
}

.page4-col {
  flex-basis: 24%;
  border-radius: 10px;
  margin-bottom: 30px;
  position: relative;
  overflow: hidden;
}

.page4-col img {
  width: 100%;
  height: 100%;
  display: block;
}

.photoj {
  height: 500px;
  margin-left: 5px;
  margin-right: 30px;
  border-radius: 40%;
}

.photopendu {
  margin-left: 50px;
  height: 500px;
  height: 900px;
}

.icontea {
  height: 80px;
}

.imgynov {
  height: 300px;
}

.photoflav {
  height: 500px;
  margin-left: 5px;
  margin-right: 30px;
  border-radius: 40%;
}

</style>