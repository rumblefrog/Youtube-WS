import { ActionTree } from 'vuex';

const { Toast } = require('buefy/dist/components/toast');

import axios from 'axios';

import { SearchState, Video, Manifest } from './types';
import { RootState } from '../types';

import router from '@/router';


export const actions: ActionTree<SearchState, RootState> = {
    async searchVideos({ commit }, keyword: string): Promise<void> {
        try {
            const res = await axios.get('/api/search?keyword=' + keyword);

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
    purgeData({ commit }): any {
        commit('purgeData');
    },
    async setVCode({ commit }, vcode: string): Promise<void> {
        commit('setVCode', vcode);

        try {
            const res = await axios.get('/api/manifest?vcode=' + vcode);

            const payload: Manifest = res && res.data;

            commit('setManifest', payload);
        } catch {
            Toast.open({
                message: 'Failed to fetch video manifest',
                type: 'is-danger',
                position: 'is-bottom',
            });
        }

        router.push('/selection');
    },
};