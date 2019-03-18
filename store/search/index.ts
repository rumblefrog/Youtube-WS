import { Module } from 'vuex';
import { actions } from './actions';
import { mutations } from './mutations';
import { SearchState } from './types';
import { RootState } from '../types';

export const state: SearchState = {
    videos: [],
    currentVCode: "",
    currentManifest: undefined,
};

const namespaced: boolean = true;

export const search: Module<SearchState, RootState> = {
    namespaced,
    state,
    actions,
    mutations,
};
