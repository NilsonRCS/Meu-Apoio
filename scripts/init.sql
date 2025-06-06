-- Criação do banco de dados MeuApoio
-- Script de inicialização para PostgreSQL

-- Extensões
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    birth_date DATE,
    phone VARCHAR(20),
    profile_image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

-- Tabela de contatos de emergência
CREATE TABLE IF NOT EXISTS emergency_contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    relationship VARCHAR(50),
    is_primary BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de áudios
CREATE TABLE IF NOT EXISTS audios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    category VARCHAR(50) NOT NULL, -- 'meditation', 'relaxing_music'
    file_url TEXT NOT NULL,
    file_size BIGINT,
    duration_seconds INTEGER,
    thumbnail_url TEXT,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

-- Tabela de favoritos dos usuários
CREATE TABLE IF NOT EXISTS user_favorites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    audio_id UUID NOT NULL REFERENCES audios(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, audio_id)
);

-- Tabela de histórico de reprodução
CREATE TABLE IF NOT EXISTS user_play_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    audio_id UUID NOT NULL REFERENCES audios(id) ON DELETE CASCADE,
    played_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completion_percentage INTEGER DEFAULT 0
);

-- Tabela de notificações
CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    message TEXT NOT NULL,
    type VARCHAR(50) NOT NULL, -- 'reminder', 'support', 'system'
    is_read BOOLEAN DEFAULT false,
    scheduled_for TIMESTAMP,
    sent_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índices para otimização
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_emergency_contacts_user_id ON emergency_contacts(user_id);
CREATE INDEX IF NOT EXISTS idx_audios_category ON audios(category);
CREATE INDEX IF NOT EXISTS idx_user_favorites_user_id ON user_favorites(user_id);
CREATE INDEX IF NOT EXISTS idx_user_play_history_user_id ON user_play_history(user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications(user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_scheduled_for ON notifications(scheduled_for);

-- Dados iniciais para testes
INSERT INTO users (username, email, password_hash, full_name) VALUES 
('admin', 'admin@meuapoio.com', '$2a$10$example.hash.here', 'Administrador'),
('usuario_teste', 'teste@meuapoio.com', '$2a$10$example.hash.here', 'Usuário de Teste')
ON CONFLICT DO NOTHING;

-- Áudios de exemplo
INSERT INTO audios (title, description, category, file_url, duration_seconds) VALUES 
('Meditação para Ansiedade', 'Uma meditação guiada de 10 minutos para reduzir a ansiedade', 'meditation', '/audio/meditacao_ansiedade.mp3', 600),
('Sons da Natureza - Chuva', 'Sons relaxantes de chuva para dormir', 'relaxing_music', '/audio/chuva_relaxante.mp3', 1800),
('Respiração Consciente', 'Exercício de respiração para momentos de stress', 'meditation', '/audio/respiracao_consciente.mp3', 300)
ON CONFLICT DO NOTHING;

COMMENT ON TABLE users IS 'Tabela principal de usuários do sistema';
COMMENT ON TABLE emergency_contacts IS 'Contatos de emergência dos usuários';
COMMENT ON TABLE audios IS 'Catálogo de áudios disponíveis';
COMMENT ON TABLE user_favorites IS 'Áudios favoritos dos usuários';
COMMENT ON TABLE user_play_history IS 'Histórico de reprodução dos usuários';
COMMENT ON TABLE notifications IS 'Sistema de notificações'; 