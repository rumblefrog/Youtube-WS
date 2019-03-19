<template>
    <section id="app" class="hero is-dark is-fullheight">
        <div class="hero-head">
            <nav class="navbar" role="navigation" aria-label="Navigation">
                <div class="container">
                    <div class="navbar-brand">
                        <router-link to="/" class="navbar-item">
                            <span>Z</span>
                        </router-link>
                    </div>
                </div>
            </nav>
        </div>
        <div class="hero-body">
            <div class="container">
                <router-view></router-view>
            </div>
        </div>
    </section>
</template>

<style lang="scss">
@import "~bulma/sass/utilities/_all";
$dark: #202125;
$dark-invert: findColorInvert($dark);

$dark-blue: #0096a7;
$dark-blue-invert: findColorInvert($dark-blue);

$primary: $dark-blue;
$primary-invert: findColorInvert($primary);

$light-blue: #0cafc1;
$light-blue-invert: findColorInvert($light-blue);

$yellow: #fdde3d;
$yellow-invert: findColorInvert($yellow);

$colors: (
    "white": ($white, $black),
    "black": ($black, $white),
    "light": ($light, $light-invert),
    "dark": ($dark, $dark-invert),
    "primary": ($primary, $primary-invert),
    "info": ($info, $info-invert),
    "success": ($success, $success-invert),
    "warning": ($warning, $warning-invert),
    "danger": ($danger, $danger-invert),
    "dark-blue": ($dark-blue, $dark-blue-invert),
    "light-blue": ($light-blue, $light-blue-invert),
    "yellow": ($yellow, $yellow-invert)
);

$background: $dark;
$link: $white;
$family-sans-serif: "Montserrat", sans-serif;
$family-primary: $family-sans-serif;

@import "~bulma/sass/base/_all";
@import "~bulma/sass/layout/hero";
@import "~bulma/sass/layout/section";
@import "~bulma/sass/grid/columns";
@import "~bulma/sass/components/media";
@import "~bulma/sass/components/navbar";
@import "~bulma/sass/elements/button";
@import "~bulma/sass/elements/box";
@import "~bulma/sass/elements/container";
@import "~bulma/sass/elements/image";
@import "~bulma/sass/elements/form";
@import "~buefy/src/scss/components/_form";
@import "./style/fonts";
</style>

<style scoped>
.navbar {
    border-radius: 0;
}
</style>