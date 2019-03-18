export interface Video {
    title: string;
    creator: string;
    vcode: string;
    timelapsed: string;
    views: string;
    description: string;
    thumbnail: string;
}

export interface SearchState {
    videos: Video[];
}