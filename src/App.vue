<template>
  <head>
    <title>STAR WARS DATA PAGE</title>
    <link
      rel="icon"
      type="image/png"
      href="https://zupimages.net/up/22/06/fd4e.png"
    />
    <link rel="stylesheet" href="style.css" />
    <link
      href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
      rel="stylesheet"
    />
  </head>

  <body>
    <div class="navbar">
      <div class="icon">
        <h2 class="logo">STAR WARS DATA</h2>
      </div>
      <div class="menu">
        <ul>
          <li><a href="#page1">QUIZZ</a></li>
          <li><a href="#page2">TEAM</a></li>
          <li><a href="#page3">DATA</a></li>
          <li><a href="#git">RESEAUX</a></li>
          <li>
            <img class="logosw" src="https://zupimages.net/up/22/06/fd4e.png" />
          </li>
        </ul>
      </div>
    </div>
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
    <section class="page4" id="page4">
      <h2>Acknowledgements</h2>
      <p>
        Thanks to our professor and mentors<br />
        <img class="imgynov" src="https://zupimages.net/up/22/05/kjk1.png" />
      </p>
    </section>
    <section class="git" id="git">
      <h1>Our Repository in GITEA<br /><span>Hangman-web</span></h1>
      <p>
        <img class="icontea" src="https://zupimages.net/up/21/49/p6fk.png" />
      </p>
      <br />
      <a
        href="https://git.ytrack.learn.ynov.com/FCORMERAIS/Hangman-Web/branches"
        >Gitea</a
      >
    </section>
  </body>
  {{ end }}
</template>

<style>
* {
  /*applique à toute la page*/
  margin: 0; /*supprime le contour blanc*/
  font-family: "Teko"; /*ou "Prompt", j'ai importé 2 polices qui me plaisaient via google front*/
}

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

.navbar {
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: space-between; /* configure un éspace entre les items*/
  position: fixed; /*pendant le scroll, la barre reste fixe*/
  padding: 0 5%;
  left: 0;
  right: 0;
  background: rgb(0, 0, 0); /*J'ai ajouté de la transparence*/
}

.logo {
  color: #ffffff;
  font-size: 45px;
}

.menu {
  width: 50%;
}

ul {
  display: flex;
  justify-content: center;
  align-items: center;
}

ul li {
  list-style: none;
  padding: 8px 20px 0 20px;
  font-size: 25px;
}

ul li::after {
  content: "";
  width: 0%;
  height: 2px;
  background: rgba(240, 225, 14, 0.671);
  display: block;
  margin: auto;
}

ul li:hover::after {
  width: 100%;
}

ul li a {
  text-decoration: none;
  color: #fff;
  text-transform: uppercase;
  font-weight: bold;
  transition: 0.4s ease-in-out;
}

ul li a:hover {
  color: rgba(240, 225, 14, 0.671);
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

.git {
  background-image: linear-gradient(rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.7)),
    url(/images/5.jpg);
  background-position: center;
  background-size: cover;
  text-align: center;
  padding: 70px 0;
}

.git h1 {
  color: #fff;
  margin-bottom: 35px;
  text-transform: capitalize;
}

.git h1 span {
  color: #b19a33;
}

.git a {
  border: 1px solid #fff;
  padding: 10px 30px;
  text-decoration: none;
  color: #fff;
}

.git a:hover {
  background: #fff;
  color: #b19a33;
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

.logosw {
  height: 100px;
}
</style>
