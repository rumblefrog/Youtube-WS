export interface Video {
    title: string;
    creator: string;
    vcode: string;
    timelapsed: string;
    views: string;
    description: string;
    thumbnail: string;
}

export interface Format {
    itag: number;
    extension: string;
    resolution: string;
    videoencoding: string;
    audioencoding: string;
    audiobitrate: number;
}

export interface Manifest {
    id: string;
    title: string;
    description: string;
    datepublished: string;
    formats: Format[];
    keywords: string[];
    author: string;
    duration: string;
}

export interface SearchState {
    videos: Video[];
    currentVCode: string;
    currentManifest: Manifest | undefined;
}