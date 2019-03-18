import { ActionTree } from 'vuex';

import { Toast } from 'buefy/dist/components/toast';

import axios from 'axios';

import { SearchState, Video } from './types';
import { RootState } from '../types';

import router from '@/router';


export const actions: ActionTree<SearchState, RootState> = {
    async searchVideos({ commit }, keyword: string): Promise<any> {
        try {
            const res = await axios.get('http://localhost:8080/api/search?keyword=' + keyword);

            const payload: Video[] = res && res.data && res.data.videos;

            for (const p of payload)
                commit('addVideo', p);

            router.push('/videos');
        } catch {
            Toast.open({
                message: 'Failed to fetch videos',
                type: 'is-danger',
                position: 'is-bottom',
            });
        }
    },
    purgeVideos({ commit }): any {
        commit('purgeVideos');
    },
};