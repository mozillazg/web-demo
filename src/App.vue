<template>
  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <!-- Brand and toggle get grouped for better mobile display -->
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="/" v-link="{path: '/'}">DEMO</a>
      </div>

      <!-- Collect the nav links, forms, and other content for toggling -->
      <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
        <ul class="nav navbar-nav">
          <li class="active" v-if="$route.name === 'pinyin'">
            <a href="/" v-link="{name: 'pinyin'}">pinyin</a>
          </li>
          <li class="" v-else>
            <a href="/" v-link="{name: 'pinyin'}">pinyin</a>
          </li>

          <li class="active" v-if="$route.name === 'unidecode'">
            <a href="/" v-link="{name: 'unidecode'}">unidecode</a>
          </li>
          <li class="" v-else>
            <a href="/" v-link="{name: 'unidecode'}">unidecode</a>
          </li>

          <li class="active" v-if="$route.name === 'slugify'">
            <a href="/" v-link="{name: 'slugify'}">slugify</a>
          </li>
          <li class="" v-else>
            <a href="/" v-link="{name: 'slugify'}">slugify</a>
          </li>

          <li class="active" v-if="$route.name === 'stpinyin'">
            <a href="/" v-link="{name: 'stpinyin'}">stpinyin</a>
          </li>
          <li class="" v-else>
            <a href="/" v-link="{name: 'stpinyin'}">stpinyin</a>
          </li>
        </ul>
      </div><!-- /.navbar-collapse -->
    </div><!-- /.container-fluid -->
  </nav>

  <!-- Page Content -->
  <div class="container">

    <!-- Page Heading -->
    <div class="">
      <h1><a href="/">{{ name }}</a></h2>
    </div>
    <!-- /.row -->

    <div >
      <!-- Projects Row -->
      <div class="row">
        <router-view></router-view>
      </div>
      <!-- /.row -->
    </div>

    <!-- Footer -->
    <footer>
      <div class="row">
        <div class="col-lg-12">
        </div>
      </div>
    <!-- /.row -->
    </footer>

  </div>
  <!-- /.container -->
</template>

<script>
export default {
}
</script>
