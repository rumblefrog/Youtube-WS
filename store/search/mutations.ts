import { MutationTree } from 'vuex';
import { SearchState, Video, Manifest } from './types';

export const mutations: MutationTree<SearchState> = {
    addVideo(state, payload: Video) {
        state.videos.push(payload);
    },
    purgeData(state) {
        state.videos = [];
        state.currentManifest = undefined;
        state.currentVCode = "";
    },
    setVCode(state, payload: string) {
        state.currentVCode = payload;
    },
    setManifest(state, payload: Manifest) {
        state.currentManifest = payload;
    },
};