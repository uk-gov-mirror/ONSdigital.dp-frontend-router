package productPage

//StubbedData ...
var stubbedData = `{
  "taxonomy": [
    {
      "uri": "/businessindustryandtrade",
      "children": [
        {
          "uri": "/businessindustryandtrade/business",
          "description": {},
          "children": [
            {
              "uri": "/businessindustryandtrade/business/activitysizeandlocation",
              "description": {
                "title": "Activity, size and location"
              },
              "type": "product_page"
            },
            {
              "uri": "/businessindustryandtrade/business/businessinnovation",
              "description": {
                "title": "Business innovation"
              },
              "type": "product_page"
            },
            {
              "uri": "/businessindustryandtrade/business/businessservices",
              "description": {
                "title": "Business services"
              },
              "type": "product_page"
            }
          ],
          "title": "Business"
        },
        {
          "uri": "/businessindustryandtrade/changestobusiness",
          "description": {},
          "children": [
            {
              "uri": "/businessindustryandtrade/changestobusiness/bankruptcyinsolvency",
              "description": {
                "title": "Bankruptcy/insolvency"
              },
              "type": "product_page"
            },
            {
              "uri": "/businessindustryandtrade/changestobusiness/businessbirthsdeathsandsurvivalrates",
              "description": {
                "title": "Business births, deaths and survival rates"
              },
              "type": "product_page"
            },
            {
              "uri": "/businessindustryandtrade/changestobusiness/mergersandacquisitions",
              "description": {
                "title": "Mergers and acquisitions"
              },
              "type": "product_page"
            }
          ],
          "title": "Changes to business"
        },
        {
          "uri": "/businessindustryandtrade/constructionindustry",
          "description": {},
          "title": "Construction industry"
        },
        {
          "uri": "/businessindustryandtrade/itandinternetindustry",
          "description": {},
          "title": "IT and internet industry"
        },
        {
          "uri": "/businessindustryandtrade/internationaltrade",
          "description": {},
          "title": "International trade"
        },
        {
          "uri": "/businessindustryandtrade/manufacturingandproductionindustry",
          "description": {},
          "title": "Manufacturing and production industry"
        },
        {
          "uri": "/businessindustryandtrade/retailindustry",
          "description": {},
          "title": "Retail industry"
        },
        {
          "uri": "/businessindustryandtrade/tourismindustry",
          "description": {},
          "title": "Tourism industry"
        }
      ],
      "title": "Business, industry and trade"
    },
    {
      "uri": "/economy",
      "children": [
        {
          "uri": "/economy/economicoutputandproductivity",
          "description": {},
          "children": [
            {
              "uri": "/economy/economicoutputandproductivity/output",
              "description": {
                "title": "Output"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/economicoutputandproductivity/productivitymeasures",
              "description": {
                "title": "Productivity measures"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/economicoutputandproductivity/publicservicesproductivity",
              "description": {
                "title": "Public services productivity"
              },
              "type": "product_page"
            }
          ],
          "title": "Economic output and productivity"
        },
        {
          "uri": "/economy/environmentalaccounts",
          "description": {},
          "title": "Environmental accounts"
        },
        {
          "uri": "/economy/governmentpublicsectorandtaxes",
          "description": {},
          "children": [
            {
              "uri": "/economy/governmentpublicsectorandtaxes/localgovernmentfinance",
              "description": {
                "title": "Local government finance"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/governmentpublicsectorandtaxes/publicsectorfinance",
              "description": {
                "title": "Public sector finance"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/governmentpublicsectorandtaxes/publicspending",
              "description": {
                "title": "Public spending"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/governmentpublicsectorandtaxes/researchanddevelopmentexpenditure",
              "description": {
                "title": "Research and development expenditure"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/governmentpublicsectorandtaxes/taxesandrevenue",
              "description": {
                "title": "Taxes and revenue"
              },
              "type": "product_page"
            }
          ],
          "title": "Government, public sector and taxes"
        },
        {
          "uri": "/economy/grossdomesticproductgdp",
          "description": {},
          "title": "Gross Domestic Product (GDP)"
        },
        {
          "uri": "/economy/grossvalueaddedgva",
          "description": {},
          "title": "Gross Value Added (GVA)"
        },
        {
          "uri": "/economy/inflationandpriceindices",
          "description": {},
          "title": "Inflation and price indices"
        },
        {
          "uri": "/economy/investmentspensionsandtrusts",
          "description": {},
          "title": "Investments, pensions and trusts"
        },
        {
          "uri": "/economy/nationalaccounts",
          "description": {},
          "children": [
            {
              "uri": "/economy/nationalaccounts/balanceofpayments",
              "description": {
                "title": "Balance of payments"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/nationalaccounts/satelliteaccounts",
              "description": {
                "title": "Satellite accounts"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/nationalaccounts/supplyandusetables",
              "description": {
                "title": "Supply and use tables"
              },
              "type": "product_page"
            },
            {
              "uri": "/economy/nationalaccounts/uksectoraccounts",
              "description": {
                "title": "UK sector accounts"
              },
              "type": "product_page"
            }
          ],
          "title": "National accounts"
        },
        {
          "uri": "/economy/regionalaccounts",
          "description": {},
          "children": [
            {
              "uri": "/economy/regionalaccounts/grossdisposablehouseholdincome",
              "description": {
                "title": "Gross disposable household income"
              },
              "type": "product_page"
            }
          ],
          "title": "Regional accounts"
        }
      ],
      "title": "Economy"
    },
    {
      "uri": "/employmentandlabourmarket",
      "children": [
        {
          "uri": "/employmentandlabourmarket/peopleinwork",
          "description": {},
          "children": [
            {
              "uri": "/employmentandlabourmarket/peopleinwork/earningsandworkinghours",
              "description": {
                "title": "Earnings and working hours"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes",
              "description": {
                "title": "Employment and employee types"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peopleinwork/labourproductivity",
              "description": {
                "title": "Labour productivity"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peopleinwork/publicsectorpersonnel",
              "description": {
                "title": "Public sector personnel"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions",
              "description": {
                "title": "Workplace disputes and working conditions"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peopleinwork/workplacepensions",
              "description": {
                "title": "Workplace pensions"
              },
              "type": "product_page"
            }
          ],
          "title": "People in work"
        },
        {
          "uri": "/employmentandlabourmarket/peoplenotinwork",
          "description": {},
          "children": [
            {
              "uri": "/employmentandlabourmarket/peoplenotinwork/economicinactivity",
              "description": {
                "title": "Economic inactivity"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peoplenotinwork/outofworkbenefits",
              "description": {
                "title": "Out of work benefits"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peoplenotinwork/redundancies",
              "description": {
                "title": "Redundancies"
              },
              "type": "product_page"
            },
            {
              "uri": "/employmentandlabourmarket/peoplenotinwork/unemployment",
              "description": {
                "title": "Unemployment"
              },
              "type": "product_page"
            }
          ],
          "title": "People not in work"
        }
      ],
      "title": "Employment and labour market"
    },
    {
      "uri": "/peoplepopulationandcommunity",
      "children": [
        {
          "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/adoption",
              "description": {
                "title": "Adoption"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/ageing",
              "description": {
                "title": "Ageing"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/conceptionandfertilityrates",
              "description": {
                "title": "Conception and fertility rates"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths",
              "description": {
                "title": "Deaths"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/divorce",
              "description": {
                "title": "Divorce"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/families",
              "description": {
                "title": "Families"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/lifeexpectancies",
              "description": {
                "title": "Life expectancies"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths",
              "description": {
                "title": "Live births"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/marriagecohabitationandcivilpartnerships",
              "description": {
                "title": "Marriage, cohabitation and civil partnerships"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/maternities",
              "description": {
                "title": "Maternities"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/birthsdeathsandmarriages/stillbirths",
              "description": {
                "title": "Stillbirths"
              },
              "type": "product_page"
            }
          ],
          "title": "Births, deaths and marriages"
        },
        {
          "uri": "/peoplepopulationandcommunity/crimeandjustice",
          "description": {},
          "title": "Crime and justice"
        },
        {
          "uri": "/peoplepopulationandcommunity/culturalidentity",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/culturalidentity/ethnicity",
              "description": {
                "title": "Ethnicity"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/culturalidentity/language",
              "description": {
                "title": "Language"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/culturalidentity/religion",
              "description": {
                "title": "Religion"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/culturalidentity/sexuality",
              "description": {
                "title": "Sexuality"
              },
              "type": "product_page"
            }
          ],
          "title": "Cultural identity"
        },
        {
          "uri": "/peoplepopulationandcommunity/educationandchildcare",
          "description": {},
          "title": "Education and childcare"
        },
        {
          "uri": "/peoplepopulationandcommunity/elections",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/elections/electoralregistration",
              "description": {
                "title": "Electoral registration"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/elections/generalelections",
              "description": {
                "title": "General elections"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/elections/localgovernmentelections",
              "description": {
                "title": "Local government elections"
              },
              "type": "product_page"
            }
          ],
          "title": "Elections"
        },
        {
          "uri": "/peoplepopulationandcommunity/healthandsocialcare",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/causesofdeath",
              "description": {
                "title": "Causes of death"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/childhealth",
              "description": {
                "title": "Child health"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases",
              "description": {
                "title": "Conditions and diseases"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/disability",
              "description": {
                "title": "Disability"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/drugusealcoholandsmoking",
              "description": {
                "title": "Drug use, alcohol and smoking"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/healthandlifeexpectancies",
              "description": {
                "title": "Health and life expectancies"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/healthandwellbeing",
              "description": {
                "title": "Health and well-being"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/healthcaresystem",
              "description": {
                "title": "Health care system"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/healthinequalities",
              "description": {
                "title": "Health inequalities"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/mentalhealth",
              "description": {
                "title": "Mental health"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/healthandsocialcare/socialcare",
              "description": {
                "title": "Social care"
              },
              "type": "product_page"
            }
          ],
          "title": "Health and social care"
        },
        {
          "uri": "/peoplepopulationandcommunity/householdcharacteristics",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/householdcharacteristics/homeinternetandsocialmediausage",
              "description": {
                "title": "Home internet and social media usage"
              },
              "type": "product_page"
            }
          ],
          "title": "Household characteristics"
        },
        {
          "uri": "/peoplepopulationandcommunity/housing",
          "description": {},
          "title": "Housing"
        },
        {
          "uri": "/peoplepopulationandcommunity/leisureandtourism",
          "description": {},
          "title": "Leisure and tourism"
        },
        {
          "uri": "/peoplepopulationandcommunity/personalandhouseholdfinances",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/personalandhouseholdfinances/debt",
              "description": {
                "title": "Debt"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/personalandhouseholdfinances/expenditure",
              "description": {
                "title": "Expenditure"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/personalandhouseholdfinances/incomeandwealth",
              "description": {
                "title": "Income and wealth"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/personalandhouseholdfinances/pensionssavingsandinvestments",
              "description": {
                "title": "Pensions, savings and investments"
              },
              "type": "product_page"
            }
          ],
          "title": "Personal and household finances"
        },
        {
          "uri": "/peoplepopulationandcommunity/populationandmigration",
          "description": {},
          "children": [
            {
              "uri": "/peoplepopulationandcommunity/populationandmigration/internationalmigration",
              "description": {
                "title": "International migration"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/populationandmigration/migrationwithintheuk",
              "description": {
                "title": "Migration within the UK"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/populationandmigration/populationestimates",
              "description": {
                "title": "Population estimates"
              },
              "type": "product_page"
            },
            {
              "uri": "/peoplepopulationandcommunity/populationandmigration/populationprojections",
              "description": {
                "title": "Population projections"
              },
              "type": "product_page"
            }
          ],
          "title": "Population and migration"
        },
        {
          "uri": "/peoplepopulationandcommunity/wellbeing",
          "description": {},
          "title": "Well-being"
        }
      ],
      "title": "People, population and community"
    }
  ],
  "uri": "/",
  "type": "product_page",
  "metadata": {
    "title": "Home",
    "description": "thing",
    "keywords": [
      "statistics",
      "economy",
      "census",
      "population",
      "inflation",
      "employment"
    ]
  },
  "data": {}
}`
