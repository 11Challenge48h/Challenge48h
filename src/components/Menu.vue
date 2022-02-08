<script setup>
import { useRouter } from 'vue-router';
import { RouterLink, RouterView } from 'vue-router';

const router = useRouter();

</script>

<template>
    <div class="navbar">
        <h2 class="logo">STAR WARS DATA</h2>
        <div class="menu">
            <li><RouterLink class="li1" to="/">ACCUEIL</RouterLink></li>
            <li><RouterLink class="li2" to="/Team">EQUIPE</RouterLink></li>
            <li><RouterLink class="li3" to="/Quiz">QUIZ</RouterLink></li> 
            <li><RouterLink class="li4" to="/Wiki">WIKI</RouterLink></li> 
            <li>
                <img class="logosw" src="https://zupimages.net/up/22/06/fd4e.png" />
            </li>
        </div>
    </div>
</template>

<style>

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
  display: flex;
  justify-content: center;
  align-items: center;
}

.logosw {
  height: 100px;
}

li {
  list-style: none;
  padding: 8px 20px 0 20px;
  font-size: 25px;
}
li::after {
  content: "";
  width: 0%;
  height: 2px;
  background: rgba(240, 225, 14, 0.671);
  display: block;
  margin: auto;
}

li:hover::after {
  width: 100%;
}

li a {
  text-decoration: none;
  color: #fff;
  text-transform: uppercase;
  font-weight: bold;
  transition: 0.4s ease-in-out;
}

li .li1:hover {
  color: rgba(240, 225, 14, 0.671);
}

li .li2:hover {
  color: rgba(240, 225, 14, 0.671);
}

li .li3:hover {
  color: rgba(240, 225, 14, 0.671);
}

li .li4:hover {
  color: rgba(240, 225, 14, 0.671);
}

</style>
