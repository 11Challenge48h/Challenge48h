<script setup>

</script>

<template>
 <head>
    <link 
    href="https://fonts.googleapis.com/css2?family=Prompt:wght@200&family=Supermercado+One&family=Teko:wght@300&display=swap"
    rel="stylesheet"
    />
    <title>STAR WARS DATA PAGE</title>
  </head>
  <body>
    <section class="main">
      <section class="page1">
        <div class="intro">
          <h1>
            <div class="Ombre">
            STAR WARS DATA
            </div>
            <span>May the Force be with you !</span>
          </h1>
          <p class="text Ombre">
            This site was made in 48H since it comes from the Challenge 48h school project Ynov Nantes. 
            We are a group of computer science students 
            and we offer you this Star Wars specific QUIZ and Wiki 
          </p>
          
        </div>

        <div>
          <img
            class="photodroid"
            src="https://zupimages.net/up/22/06/ns5k.png"
          />
        </div>
      </section>
    </section>
    <section class="page2"> <!-- MEttre les boutons ici -->
      <RouterLink to="/quizz"><button>PLAY NOW</button></RouterLink>
      <RouterLink to="/wiki"><button>WIKI</button></RouterLink>
    </section>
    <section class="page3">
      <br/>
      <br/>
      <h2>Pour en savoir plus sur notre équipe ....</h2>
    </section>
  </body>
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

.Ombre {
  text-shadow: 2px 2px 4px #000000;
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

.page2 > button {
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
  margin-top: 10px;
}

.intro > button:hover {
  color: #000000;
  background-color: white;
}

.intro span {
  color: #000000;
  font-size: 65px;
}

.page2,
.page3 {
  border-radius: 5px;
  padding: 10% 15%;
}

.page3 {
  background: #753e1914;
}

.page2 h2,
.page3 h2 {
  margin-bottom: 25px;
  font-size: 30px;
  font-weight: bold;
  color: #000;
  text-align: center;
}

.page2 p,
.page3 p {
  text-align: center;
}

.photodroid {
  margin-left: 50px;
  height: 500px;
  height: 900px;
}

</style>