import React from 'react';

const VideoPlayer = ({ param }) => {
    const title = param.title
    const quality = param.quality
    const videoSrc = `${process.env.NEXT_PUBLIC_API_BASE_URL}/streamer/stream?title=${title}&quality=${quality}`;

    return (
        <video controls width="750">
            <source src={videoSrc} type="video/mp4" />
            Your browser does not support the video tag.
        </video>
    );
};

export default VideoPlayer;
