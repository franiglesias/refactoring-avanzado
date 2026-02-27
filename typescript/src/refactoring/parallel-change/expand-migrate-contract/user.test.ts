import {describe, expect, it} from 'vitest'
import {buildUserSummary, formatDisplayName, formatEmailHeader, formatGreeting, User} from './user'

describe('Expand-Migrate-Contract (Parallel Change)', () => {
  const alice = new User('u-1', 'Alice Smith', 'alice@example.com')
  const bob = new User('u-2', 'Bob Jones', 'bob@example.com')

  describe('formatGreeting', () => {
    it('should greet the user by name', () => {
      expect(formatGreeting(alice)).toBe('Hello, Alice Smith!')
    })
  })

  describe('formatEmailHeader', () => {
    it('should format the email header with name and email', () => {
      expect(formatEmailHeader(alice)).toBe('From: Alice Smith <alice@example.com>')
    })
  })

  describe('formatDisplayName', () => {
    it('should format name with id', () => {
      expect(formatDisplayName(alice)).toBe('Alice Smith (u-1)')
    })
  })

  describe('buildUserSummary', () => {
    it('should list all user names', () => {
      expect(buildUserSummary([alice, bob])).toBe('- Alice Smith\n- Bob Jones')
    })

    it('should return empty string for empty list', () => {
      expect(buildUserSummary([])).toBe('')
    })
  })
})
