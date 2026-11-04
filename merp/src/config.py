from enum import Enum


class CharName(Enum):
    STRENGTH = "strength"
    AGILITY = "agility"
    CONSTITUTION = "constitution"
    INTELLIGENCE = "intelligence"
    INTITUTION = "intuition"
    PRESENCE = "presence"

class Profession(Enum):
    WARRIOR = "warrior"
    SCOUT = "scout"
    RANGER = "ranger"
    MAGE = "mage"
    ANIMIST = "animist"
    BARD = "bard"

LOW_ROLL = 1
HIGH_ROLL = 102

MAX_SKILL_RANKS = 999