<template>
    <div id="Selection" class="columns">
        <div class="column is-4 is-offset-4">
            <div class="box"
            v-for="f in search.currentManifest.formats"
            v-bind:key="f.itag"
            @click="download(search.currentVCode, f.itag)">
                <p>
                    Resolution - {{ formatField(f.resolution) }}
                </p>

                <br />

                <p>
                    Video Encoding - {{ formatField(f.videoEncoding) }}
                </p>

                <br />

                <p>
                    Audio Bitrate - {{ formatField(f.audioBitrate) }}
                </p>

                <br />

                <p>
                    Audio Encoding - {{ formatField(f.audioEncoding) }}
                </p>

                <br />

                <p>
                    Extension - {{ formatField(f.extension) }}
                </p>

            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue} from 'vue-property-decorator';
import { State } from 'vuex-class';
import { SearchState } from '../store/search/types';

const namespace: string = 'search';

@Component
export default class Selection extends Vue {
    @State('search') public search: SearchState;

    public formatField(rate: string) {
        if (rate === '') {
            return 'N/A';
        }

        return rate;
    }

    public download(id: string, itag: number) {
        window.open('/api/video?vcode=' + id + '&itag=' + itag);
    }
}
</script>
