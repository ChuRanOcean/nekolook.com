import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [videos, setVideos] = useState([]);

  useEffect(() => {
    // 模拟获取视频数据
    setVideos([
      { id: 1, title: "视频1", description: "描述1" },
      { id: 2, title: "视频2", description: "描述2" },
      // 你可以根据需要添加更多视频数据
    ]);
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <h1>NekoLook 喵看</h1> {/* 修改标题为NekoLook 喵看 */}
      </header>
      <main>
        <div className="video-grid">
          {videos.map((video) => (
            <div key={video.id} className="video-card">
              <img 
                className="video-thumbnail" 
                src={`https://placeimg.com/400/225/tech?${video.id}`} 
                alt={video.title}
              />
              <div className="video-info">
                <h3>{video.title}</h3>
                <p>{video.description}</p>
              </div>
            </div>
          ))}
        </div>
      </main>
      <footer className="App-footer">
        <p>&copy; 2025 NekoLook 喵看. All rights reserved.</p>
      </footer>
    </div>
  );
}

export default App;
