<script setup>
  import {reactive, onMounted, ref} from 'vue'

  let mix = reactive( {
    id: 0,
    ytid: "",
    title: ""
  })

  const url = 'http://localhost:8080/api/v1/mix'
  let embedURL = ref("")

  const getMix = async () => {
    
    try {
      const response = await fetch(url)
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`)
      }
      const json = await response.json()
      mix.title = json.title
      mix.ytid = json.ytid
      mix.id = json.id
      embedURL.value = `https://www.youtube.com/embed/${mix.ytid}?autoplay=1`

    } catch (error) {
      console.error(error.message)
    }
  }
  onMounted(() => {
    getMix()
  })

</script>

<template>
  <div class="container">
    <img class="logo-main" src="/SVG/Cueclub_Logo_Bordered_01.svg" alt="">

    <iframe
      v-if="embedURL"
      class="iframe"
      :src="embedURL"
      title="YouTube video player"
      frameborder="0"
      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
      referrerpolicy="strict-origin-when-cross-origin"
      allowfullscreen
    />
    <h1>{{ mix.title }}</h1>
    <button @click = "getMix">shuffle.</button>
  </div>
</template>

<style scoped>
  h1{
    font-size: xx-large;
  }
  .container{
    display: grid;
    padding: 24px;
    gap: 16px;
    max-width: 1536px;
  }
  .iframe {
    margin-block: 64px;
    aspect-ratio: 16/9;
    width: 100%;
    border-radius: 16px;
    border: solid 2px #5e7065;
  }
  .logo-main{
    width: 128px;
    padding: 16px;
  }
  button{
    padding: 12px;
    border-radius: 8px;
    background-color: #161c18;
    border: solid 2px #5e7065;
    color: #d0ddda;
    font-weight: 700;
    font-size: medium;
    letter-spacing: .05rem;
    width: fit-content;
    transition: ease-in-out;
    transition-duration: 400ms;
    
  }
  button:hover{
    background-color:  #253028;
    transition: ease-in-out;
    transition-duration: 400ms;
  }
</style>