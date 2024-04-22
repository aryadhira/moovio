/** @type {import('next').NextConfig} */
const nextConfig = {
    images: {
        remotePatterns: [
            {
                hostname: "yts.mx"
            }
        ]
    }
};

export default nextConfig;
