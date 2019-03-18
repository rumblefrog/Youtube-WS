<template>
    <div class="box">
        <article class="media">
            <figure class="media-left">
                <p class="image">
                    <img :src="thumbnail">
                </p>
            </figure>
            <div class="media-content">
                <h2>{{ title }}</h2>

                <br />

                <h4>By <small>{{ author }}</small></h4>

                <br />

                <h4> {{ views }} - {{ timelapsed }}</h4>
            </div>
        </article>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class Video extends Vue {
    @Prop() public thumbnail!: string;
    @Prop() public title!: string;
    @Prop() public author!: string;
    @Prop() public views!: string;
    @Prop() public timelapsed!: string;
}
</script>
