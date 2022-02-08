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
    <div class="optionaladd">
      <button type="title" title="Ynov" onclick="window.location.href='https://www.ynov.com/';" id="button-sticky"></button>
    <section class="main">
      <section class="page1">
        <div class="intro">
          <h1>
            <div class="Ombre">
            STAR WARS DATA
            </div>
            <span>May the Force be with you !</span>
          </h1>
          <p class="textOmbre">
            This site was made in 48H since it comes from the Challenge 48h school project Ynov Nantes. <br>
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
     <RouterLink tag="li" class="button1" to="/quizz"><button>PLAY NOW</button></RouterLink>
     <RouterLink tag="li" class="button2" to="/wiki"><button>WIKI</button></RouterLink>
     <div class="textUnderButton">
      <p>
       You can choose between these two possibilities.
The Wiki presents all the data of the Star Wars universe in the form of <br> 
categories such as planets, species, characters, etc... The Quizz allows the discovery of the Star Wars universe in a lucrative way.
      </p>
     </div>
    </section>
    <section class="page3">
      <br/>
      <br/>
      <h2>Pour en savoir plus sur notre équipe ....</h2>
      <RouterLink tag="li" class="buttonTeam" to="/Team"><button>Team Challenge 48h</button></RouterLink>
    </section>
    </div>
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
  height: 75vh; /*configure pour que la première page remplisse l'écran*/
  display: flex;
}
.page1 {
  color: #fff;
  display: flex;
  margin: 20% 0;
}
.optionaladd #button-sticky {
  position:fixed;
  bottom:3%;
  right:2%;
  z-index:9999;
}

.optionaladd #button-sticky:hover {
  opacity: 0.8;
  border: 1px solid #000;
  border-radius: 30px;
}
.Ombre {
  text-shadow: 3px 3px 7px #000000;
  font-size: 70px;
}

.textOmbre{
  color: #000;
  font-size: 20px;
}

.textUnderButton {
  margin-top: 2%; 
  margin-left: -560px;
  font-size: 20px;
}
.page1 .text {
  padding-bottom: 100px;
  letter-spacing: 1.2px;
  line-height: 30px;
}
.intro h1 {
  font-size: 50px;
  letter-spacing: 2px;
  margin-bottom: 35px;
}
.page2 .button1 button{
  margin-left: 5%;
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

.page2 .button2 button{
  margin-left: 3%;
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
.page2 .button2 button:hover {
  color: white;
  background-color: #000000;
  opacity: 0.8;
}
.page2 .button1 button:hover {
  color: white;
  background-color: #000000;
  opacity: 0.8;
}
.intro span {
  color: #000000;
  font-size: 35px;
  font-family: 'Trebuchet MS', Helvetica, sans-serif;
}
.page2,
.page3 {
  border-radius: 5px;
  padding: 10% 15%;
}
.page3 {
background-color: rgb(179, 174, 174)
}

.page3 .buttonTeam button{
  margin-left: 5%;
  width: 200px;
  height: 50px;
  border: 4.5px solid rgb(0, 0, 0);
  background-color: transparent;
  font-size: 30px;
  cursor: pointer;
  transition: 0.8s ease;
  color: rgb(0, 0, 0);
  /* border-radius: 10%; */
  margin-top: -1%;
  margin-left: 40%;
}

.page3 .buttonTeam button:hover {
  color: white;
  background-color: #000000;
  opacity: 0.8;
  text-shadow: 3px 3px 7px #beb5b5;
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
  height: 900px;
}
</style>
