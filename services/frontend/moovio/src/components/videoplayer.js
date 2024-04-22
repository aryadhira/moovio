import React from 'react';

const VideoPlayer = ({ param }) => {
    const title = param.title
    const quality = param.quality
    const videoSrc = `http://localhost:9003/stream?title=${title}&quality=${quality}`;

    return (
        <video controls width="750">
            <source src={videoSrc} type="video/mp4" />
            Your browser does not support the video tag.
        </video>
    );
};

export default VideoPlayer;
