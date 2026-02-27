from datetime import datetime


class UserAccount:
    """Esta clase hace demasiadas cosas: autenticación, autorización,
    gestión de perfil, notificaciones, logs, etc.
    """

    def __init__(self, user_id: str, username: str, email: str, password_hash: str):
        self.user_id = user_id
        self.username = username
        self.email = email
        self.password_hash = password_hash
        self.is_active = True
        self.roles: list[str] = []
        self.last_login: datetime = None
        self.failed_login_attempts = 0
        self.preferences: dict = {}

    def authenticate(self, password: str) -> bool:
        """Autenticación"""
        if self.check_password(password):
            self.last_login = datetime.now()
            self.failed_login_attempts = 0
            self.log_event('LOGIN_SUCCESS')
            return True
        else:
            self.failed_login_attempts += 1
            self.log_event('LOGIN_FAILED')
            if self.failed_login_attempts > 3:
                self.lock_account()
            return False

    def check_password(self, password: str) -> bool:
        """Verificar contraseña (simulado)"""
        return self.password_hash == f"hash_{password}"

    def lock_account(self):
        """Bloquear cuenta"""
        self.is_active = False
        self.log_event('ACCOUNT_LOCKED')
        self.send_notification('Tu cuenta ha sido bloqueada')

    def unlock_account(self):
        """Desbloquear cuenta"""
        self.is_active = True
        self.failed_login_attempts = 0
        self.log_event('ACCOUNT_UNLOCKED')

    def has_permission(self, permission: str) -> bool:
        """Verificar permisos"""
        return permission in self.roles

    def add_role(self, role: str):
        """Agregar rol"""
        if role not in self.roles:
            self.roles.append(role)
            self.log_event(f'ROLE_ADDED:{role}')

    def remove_role(self, role: str):
        """Remover rol"""
        if role in self.roles:
            self.roles.remove(role)
            self.log_event(f'ROLE_REMOVED:{role}')

    def update_profile(self, username: str = None, email: str = None):
        """Actualizar perfil"""
        if username:
            self.username = username
        if email:
            self.email = email
        self.log_event('PROFILE_UPDATED')
        self.send_notification('Tu perfil ha sido actualizado')

    def change_password(self, old_password: str, new_password: str) -> bool:
        """Cambiar contraseña"""
        if not self.check_password(old_password):
            return False
        self.password_hash = f"hash_{new_password}"
        self.log_event('PASSWORD_CHANGED')
        self.send_notification('Tu contraseña ha sido cambiada')
        return True

    def set_preference(self, key: str, value: any):
        """Establecer preferencia"""
        self.preferences[key] = value
        self.log_event(f'PREFERENCE_SET:{key}')

    def get_preference(self, key: str, default=None):
        """Obtener preferencia"""
        return self.preferences.get(key, default)

    def send_notification(self, message: str):
        """Enviar notificación"""
        print(f"[NOTIF] Enviando a {self.email}: {message}")

    def log_event(self, event: str):
        """Registrar evento"""
        print(f"[LOG] {datetime.now().isoformat()} - User {self.user_id}: {event}")

    def delete_account(self):
        """Eliminar cuenta"""
        self.is_active = False
        self.log_event('ACCOUNT_DELETED')
        self.send_notification('Tu cuenta ha sido eliminada')
