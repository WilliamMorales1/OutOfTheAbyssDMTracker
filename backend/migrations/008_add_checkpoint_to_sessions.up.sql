ALTER TABLE Sessions ADD COLUMN IF NOT EXISTS checkpoint TEXT;

UPDATE Sessions SET checkpoint = CASE session_num
  WHEN 1  THEN 'Party has escaped Velkynvelve and is fleeing into the Underdark with their fellow prisoners. Ilvara''s pursuit has begun.'
  WHEN 2  THEN 'Party survived the Silken Paths and is traveling deeper into the Underdark while drow pursuit escalates behind them.'
  WHEN 3  THEN 'Party escaped Sloobludop by boat after witnessing Demogorgon rise from the Darklake and destroy the settlement.'
  WHEN 4  THEN 'Party crossed the Darklake and passed through Gracklstugh''s gates into the City of Blades.'
  WHEN 5  THEN 'Party is inside Gracklstugh investigating the demonic madness and has made contact with Errde Blackskull.'
  WHEN 6  THEN 'Party reached Neverlight Grove with Stool and has spoken to at least one myconid sovereign.'
  WHEN 7  THEN 'Party is inside Blingdenstone and has been tasked by the deep gnomes with clearing the ooze infestation.'
  WHEN 8  THEN 'Party escaped the Underdark, defeated Ilvara, and received their mission briefing from the Council of Waterdeep.'
  WHEN 9  THEN 'Party returned to the Underdark and arrived at Mantol-Derith, navigating the four-faction war.'
  WHEN 10 THEN 'Party reached Araj and struck the bargain with Vizeran DeVir, receiving the full component list for the Dark Heart ritual.'
  WHEN 11 THEN 'Party entered the Wormwrithings and is actively hunting Yeenoghu''s heart as a ritual component.'
  WHEN 12 THEN 'Party is inside the Labyrinth and has located the Maze Engine''s antechamber with Baphomet closing in.'
  WHEN 13 THEN 'Party secured components from Juiblex and Zuggtmoy and has discovered Vizeran''s true plan to destroy Menzoberranzan.'
  WHEN 14 THEN 'Party infiltrated Menzoberranzan and is positioning the ritual components at Narbondellyn''s center.'
  WHEN 15 THEN 'Party is traveling through Araumycos with all components secured and the ritual finalized.'
  WHEN 16 THEN 'Party stands at Menzoberranzan''s heart with every component in place and Vizeran ready to begin the ritual.'
  ELSE NULL
END;
