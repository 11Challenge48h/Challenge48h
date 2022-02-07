<script setup>
import { useRouter } from 'vue-router';
import { RouterLink, RouterView } from 'vue-router';

const router = useRouter();

</script>

<template>
    <div class="navbar">
        <h2 class="logo">STAR WARS DATA</h2>
        <div class="menu">
            <ul>
            <li><RouterLink to="/">ACCUEIL</RouterLink></li>
            <li><RouterLink to="/Team">EQUIPE</RouterLink></li>
            <!--li><RouterLink to="Quiz">QUIZ</RouterLink></li> 
            <li><RouterLink to="Wiki">WIKI</RouterLink></li--> 
            <ul>
                <img class="logosw" src="https://zupimages.net/up/22/06/fd4e.png" />
            </ul>
            <RouterView></RouterView>
            </ul>
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

li ul:hover {
  color: rgba(240, 225, 14, 0.671);
}

ul {

}
</style>