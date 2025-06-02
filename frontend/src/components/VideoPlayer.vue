<script setup>
  import {reactive, onMounted, ref} from 'vue'

  let mix = reactive( {
    id: 0,
    ytid: "",
    title: ""
  })

  const url = 'http://localhost:8080/api/v1/mix'
  let embedURL = ref("")
  const c = ref(null)

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
    const canvas = c.value
    const ctx = canvas.getContext("2d")
    canvas.width=1536;
    canvas.height=864;

    let hue = 0
    function animate() {
      ctx.fillStyle = `hsl(${hue}, 80%, 90%)`;
      ctx.fillRect(0,0, canvas.width, canvas.height)
      ctx.effe
      hue = (hue + 1) % 360
      requestAnimationFrame(animate)
    }
    animate()
  })

</script>

<template>
  <div class="container">
    <div class="top">
      <img class="logo-main" src="/SVG/Cueclub_Logo_Bordered_01.svg" alt="">
      <button @click = "getMix">shuffle.</button>
    </div>
    <div class="vid-container">
      <canvas ref="c"></canvas>
      <div class="vid-wrapper">
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
      </div>
    </div>
    <h1>{{ mix.title }}</h1>
  </div>
</template>

<style scoped>
  h1{
    font-size: 48px;
    font-weight: 900;
    color: #99aca0;
  }
  .top{
    display: flex;
    justify-content: space-between;
    align-items: center;

  }
  .container{
    margin: 0 auto;
    display: grid;
    padding: 24px;
    gap: 16px;
    max-width: 1536px;
  }
  .vid-container{
    margin-block: 24px;
    position: relative;
  }
  .vid-wrapper{
    position: relative;
    width: 100%;
    padding-bottom: 56.25%; /* 16:9 aspect ratio (9/16 = 0.5625) */
    height: 0;
    overflow: hidden;
    border-radius: 16px;
    border: solid 2px #5e7065;
  }
  .iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border: 0;
  }
  .logo-main{
    width: 128px;
    padding: 16px;
  }
  button{
    font-family: "Outfit";
    padding: 32px 24px;
    border-radius: 12px;
    background-color: #161c18;
    border: solid 2px #5e7065;
    color: #d0ddda;
    font-weight: 800;
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
  canvas {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0;
    filter: blur(16px);
  }
</style>