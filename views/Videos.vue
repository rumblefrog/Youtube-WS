<template>
    <div id="Videos" class="columns">
        <div class="column is-6 is-offset-3">
            <Video v-for="v in search.videos"
                v-bind:key="v.vcode"
                :thumbnail="v.thumbnail"
                :title="v.title"
                :author="v.creator"
                :views="v.views"
                :timelapsed="v.timelapsed">
            </Video>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue} from 'vue-property-decorator';
import { State, Action, Getter } from 'vuex-class';
import { SearchState } from '../store/search/types';
import Video from '../components/Video.vue';

// const namespace: string = 'search';

@Component({
    components: {
        Video,
    },
})
export default class Videos extends Vue {
    @State('search') search: SearchState;
    
}
</script>
