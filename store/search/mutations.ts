import { MutationTree } from 'vuex';
import { SearchState, Video } from './types';

export const mutations: MutationTree<SearchState> = {
    addVideo(state, payload: Video) {
        state.videos.push(payload);
    },
    purgeVideos(state) {
        state.videos = [];
    },
};