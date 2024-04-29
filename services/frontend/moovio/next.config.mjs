/** @type {import('next').NextConfig} */
const nextConfig = {
    output: 'standalone',
    images: {
        remotePatterns: [
            {
                hostname: "yts.mx"
            }
        ]
    }
};

export default nextConfig;
